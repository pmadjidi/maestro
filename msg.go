package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	. "maestro/api"
	"sync"
	"time"
)

type msgEnvelope struct {
	req  chan *MsgReq
	resp chan *MsgResp
}


type Message struct {
	*MsgReq
	*sync.RWMutex
	*Flag
}

func newMessage(m *MsgReq) *Message {
	msg := Message{m,&sync.RWMutex{},NewFlag()}
	msg.Id = uuid.New().String()
	return &msg
}

type messagesdb struct {
	msg map[string]map[string]*Message
	subscriptions map[string]*User
	dirty []*Message
	blocked []*Message
	dirtyCounter int64
	blockedCounter int64
}

func newMessageDb () *messagesdb {
	return &messagesdb{make(map[string]map[string]*Message),
		make(map[string]*User),
		make([]*Message,0),make([]*Message,0),0,0}
}




func newMsgEnvelope() *msgEnvelope {
	return &msgEnvelope{make(chan *MsgReq, 1), make(chan *MsgResp, 1)}
}

type msgService struct {
	name  string
	Q     chan *msgEnvelope
	cfg   *ServerConfig
	stats *metrics
}

func (m *msgService) Name() string {
	return "msgService"
}

func newMsgService(Q chan *msgEnvelope, cfg *ServerConfig) *msgService {
	return &msgService{"msgService", Q, cfg, newMetrics()}
}

func (m *msgService) Msg(srv Message_MsgServer) error {
	log.Println("start new server")
	ctx := srv.Context()
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
			log.Printf("receive error %v", err)
			continue
		}
		//validate req
		//get and id for the request

		fmt.Printf("Got message: %s",req.Id)

		e := newMsgEnvelope()

		e.req <- req

		select {
		case m.Q <- e:
		case <-time.After(time.Second):
			srv.Send(&MsgResp{Id: "", Status: Status_TIMEOUT})
		case <-ctx.Done():
			return ctx.Err()
		}

		select {
		case resp := <-e.resp:
			if err := srv.Send(resp); err != nil {
				log.Printf("send error %v", err)
				return err
			} else {
				log.Printf("sending response =%s", resp.GetId())
			}
		case <-time.After(time.Second):
			return fmt.Errorf("Connection with grpc client is broken, timeout...")
		}
	}
}


