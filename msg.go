package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	. "maestro/api"
	"sync"
	"time"
)

type notify struct{}

type msgEnvelope struct {
	messages []*Message
	resp     chan notify
	userName string
	Status
}

type Message struct {
	id string
	*MsgReq
	*sync.RWMutex
	*Flag
}

func newMessage(msgreq *MsgReq) *Message {
	m := &Message{uuid.New().String(), msgreq, &sync.RWMutex{}, NewFlag()}
	return m
}


type messagesdb struct {
	msg             map[string]map[string]*Message
	topics  			map[*Topic]Status
	subscriptions   map[string][]*User
	mdirty          []*Message
	mblocked        []*Message
	mdirtyCounter   int64
	mblockedCounter int64
}

func newMessageDb() *messagesdb {
	return &messagesdb{make(map[string]map[string]*Message),
		make(map[*Topic]Status),
		make(map[string][]*User),
		make([]*Message, 0), make([]*Message, 0), 0, 0}
}

func newMsgEnvelope() *msgEnvelope {
	return &msgEnvelope{nil, make(chan notify,1), "", Status_NEW}
}

type msgService struct {
	name   string
	stats  *metrics
	system *Server
}

func (m *msgService) getname() string {
	return "msgService"
}

func newMsgService(s *Server) *msgService {
	return &msgService{"msgService", newMetrics(), s}
}

func (m *msgService) Put(srv Msg_PutServer) error {
	m.system.log("Called put...")

	var appName []string
	ctx := srv.Context()

	md, val := metadata.FromIncomingContext(ctx)
	PrettyPrint(md)
	if val {
		appName = md.Get("app")
		if len(appName) < 1 && appName[0] != "" {
			return fmt.Errorf(Status_INVALID_APPNAME.String())
		}
	} else {
		return fmt.Errorf(Status_NOAUTH.String())
	}

	app, err := m.system.GetOrCreateApp(appName[0], false)
	if err != nil {
		m.system.log(err.Error())
		return err
	}

	m.system.log("Before loop")
	for {

		// exit if context is done
		// or continue

		select {
		case <-ctx.Done():
			m.system.log("CTX timeout")
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			m.system.log("Put got EOF")
			return nil
		}

		if err != nil {
			m.system.log(fmt.Sprintf("receive error [%s]", err.Error()))
			return err
		}

		m.system.log(fmt.Sprintf("Recived message [%s][%s]", req.Uuid, req.Topic))
		newMsg := newMessage(req)
		e := newMsgEnvelope()
		e.messages = append(e.messages, newMsg)

		select {
		case app.msgQ <- e:
		default:
		}

		select {
		case <-e.resp:
			if err := srv.Send(&MsgResp{Status: e.Status, Uuid: e.messages[0].Uuid}); err != nil {
				return fmt.Errorf(Status_ERROR.String())
			}
		case <-time.After(m.system.cfg.SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT):
			if err := srv.Send(&MsgResp{Status: Status_TIMEOUT, Uuid: e.messages[0].Uuid}); err != nil {
				log.Printf("send error %v", err)
			}
		}
	}

	m.system.log("Returning from put")

	return nil

}
