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
	messages chan []*Message
	resp     chan notify
	userName string
	Status
}

type Message struct {
	*MsgReq
	*sync.RWMutex
	*Flag
}

func newMessage(reqs ...*MsgReq) []*Message {
	res := make([]*Message, len(reqs))

	for _, msg := range reqs {
		m := Message{msg, &sync.RWMutex{}, NewFlag()}
		m.Id = uuid.New().String()
		res = append(res, &m)
	}
	return res
}

type messagesdb struct {
	msg            map[string][]*Message
	subscriptions  map[string][]*User
	dirty          []*Message
	blocked        []*Message
	dirtyCounter   int64
	blockedCounter int64
}

func newMessageDb() *messagesdb {
	return &messagesdb{make(map[string][]*Message),
		make(map[string][]*User),
		make([]*Message, 0), make([]*Message, 0), 0, 0}
}

func newMsgEnvelope() *msgEnvelope {
	return &msgEnvelope{make(chan []*Message, 1), make(chan notify), "", Status_NEW}
}

type msgService struct {
	name   string
	stats  *metrics
	system *Server
}

func (m *msgService) Name() string {
	return "msgService"
}

func newMsgService(s *Server) *msgService {
	return &msgService{"msgService", newMetrics(), s}
}

func (m *msgService) Msg(srv Message_MsgServer) error {
	//	log.Println("start new server")

	var token, appName []string
	ctx := srv.Context()

	md, val := metadata.FromIncomingContext(ctx)
	if val {
		token = md.Get("bearer-bin")
		appName = md.Get("app")
		fmt.Printf("Token := %s", token)
		if len(token) == 0 {
			return fmt.Errorf(Status_NOAUTH.String())
		}
	}

	if len(appName) != 0 && appName[0] != "" {
		return fmt.Errorf(Status_INVALID_APPNAME.String())
	}

	app, err := m.system.GetApp(appName[0])
	if err != nil {
		return err
	}

	var recievedSoFar = make([]*MsgReq,0)
	var wg sync.WaitGroup

	wg.Add(1)
	defer wg.Wait()

	go func() {
		defer wg.Done()


	}()

	loop:
	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

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

		recievedSoFar = append(recievedSoFar, req)
	}

	//validate req
	//get and id for the request

	//fmt.Printf("Got message: %s",req.Id)

	e := newMsgEnvelope()
	e.messages <- newMessage(recievedSoFar...)

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

