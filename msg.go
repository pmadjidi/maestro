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
	messages  []*Message
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
	m := &Message{uuid.New().String(),msgreq, &sync.RWMutex{}, NewFlag()}
	return m
}

type messagesdb struct {
	msg            map[string][]*Message
	subscriptions  map[string][]*User
	mdirty          []*Message
	mblocked        []*Message
	mdirtyCounter   int64
	mblockedCounter int64
}

func newMessageDb() *messagesdb {
	return &messagesdb{make(map[string][]*Message),
		make(map[string][]*User),
		make([]*Message, 0), make([]*Message, 0), 0, 0}
}

func newMsgEnvelope() *msgEnvelope {
	return &msgEnvelope{nil, make(chan notify), "", Status_NEW}
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

func (m *msgService) Put(srv Msg_PutServer ) error {
	//	log.Println("start new server")

	var appName []string
	ctx := srv.Context()

	md, val := metadata.FromIncomingContext(ctx)
	PrettyPrint(md)
	if val {
		appName = md.Get("app")
		PrettyPrint(appName)
		if len(appName) < 1 && appName[0] != "" {
			return fmt.Errorf(Status_INVALID_APPNAME.String())
		}
	} else {
		return fmt.Errorf(Status_NOAUTH.String())
	}

	app, err := m.system.GetOrCreateApp(appName[0], false)
	if err != nil {
		return err
	}

loop:
	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// receive data from stream
			req, err := srv.Recv()
			if err == io.EOF {
				// return will close stream from server side
				log.Println("exit")
				break loop
			}
			if err != nil {
				//log.Printf("MSG receive error %v", err)
				return err
			}

			m := newMessage(req)
			e := newMsgEnvelope()

			e.messages = append(e.messages, m)

			select {
			case app.msgRecQ <- e:
			case <-time.After(time.Second):
				return fmt.Errorf(Status_TIMEOUT.String())
			case <-ctx.Done():
				return ctx.Err()
			}

			select {
			case <-e.resp:
				if e.Status == Status_SUCCESS {
					return nil
				} else {
					return fmt.Errorf(e.Status.String())
				}
			case <-time.After(time.Second):
				return fmt.Errorf("Connection with grpc client is broken, timeout...")
			}

		}
	}

	return nil
}

