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

type msgEnvelope struct {
	req  chan *MsgReq
	resp  chan struct {}
	Status
}


type Message struct {
	*MsgReq
	*sync.RWMutex
	*Flag
}

func newMessage(req ...*MsgReq) *Message {
	var m *MsgReq
	if len(req) == 0 {
		m = &MsgReq{}
	} else {
		m = req[0]
	}
	msg := Message{m,&sync.RWMutex{},NewFlag()}
	msg.Id = uuid.New().String()
	return &msg
}

type messagesdb struct {
	msg map[string][]*Message
	subscriptions map[string][]*User
	dirty []*Message
	blocked []*Message
	dirtyCounter int64
	blockedCounter int64
}

func newMessageDb () *messagesdb {
	return &messagesdb{make(map[string][]*Message),
		make(map[string][]*User),
		make([]*Message,0),make([]*Message,0),0,0}
}




func newMsgEnvelope() *msgEnvelope {
	return &msgEnvelope{make(chan *MsgReq, 1), make(chan struct{}),Status_NEW}
}

type msgService struct {
	name  string
	stats *metrics
	system *Server
}

func (m *msgService) Name() string {
	return "msgService"
}

func newMsgService(s *Server) *msgService {
	return &msgService{"msgService", newMetrics(),s}
}

func (m *msgService) Msg(srv Message_MsgServer) error {
//	log.Println("start new server")

	var token,appName []string
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

	app,err  := m.system.GetApp(appName[0])
	if err != nil {
		return err
	}



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

			return nil
		}
		if err != nil {
			//log.Printf("MSG receive error %v", err)
			return err
		}
		//validate req
		//get and id for the request

		//fmt.Printf("Got message: %s",req.Id)

		e := newMsgEnvelope()

		e.req <- req

		select {
		case app.msgQ <- e:
		case <-time.After(time.Second):
			return fmt.Errorf(Status_TIMEOUT.String())
		case <-ctx.Done():
			return ctx.Err()
		}

		select {
		case  <-e.resp:
			if e.Status == Status_SUCCESS {
				return nil
			} else {
				return  fmt.Errorf(e.Status.String())
		}
		case <-time.After(time.Second):
			return fmt.Errorf("Connection with grpc client is broken, timeout...")
		}
	}
}


