package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
	. "maestro/api"
//	"time"
	"context"
)


func registerArandomMessage(stream Message_MsgClient,numberOfMessages int) error  {
	var messagePipe = make(chan *MsgReq,numberOfMessages)
	var sendFail,reciveFail chan *error
	sendFail = make(chan *error,numberOfMessages)
	reciveFail = make(chan *error,numberOfMessages)

	go func()  {
		for msgReq := range messagePipe {
			fmt.Printf("Got message %s",msgReq.GetId())
			err := stream.SendMsg(msgReq)
			if err != nil {
				sendFail <- &err
				break
			}
		}
		close(sendFail)
	}()
	go func()  {
		for i := 0; i < numberOfMessages; i++ {
			err := stream.RecvMsg(&MsgResp{})
			if err != nil {
				reciveFail <- &err
				break
			}
		}
		close(reciveFail)
	}()

	for i := 0; i < numberOfMessages; i++ {
		r := randomMessageForTest(10)
		PrettyPrint(*r)
		messagePipe <- r
	}

	close(messagePipe)


	for e := range sendFail {
		if e != nil {
			return *e
		}
	}

	for e := range reciveFail {
		if e != nil {
			return *e
		}
	}

	return nil

}


func Test_Msg(t *testing.T) {
	// Set up a connection to the server.
	NUMBER_OF_TRAIL := 5
	//clientDeadline := time.Now().Add(time.Duration(20) * time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	//defer cancel()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewMessageClient(conn)


	for i := 0; i < NUMBER_OF_TRAIL; i++ {
		stream,err := c.Msg(context.Background())
		if err != nil {
			t.Errorf("Should not fail rpc... %+v", err)
		}  else {
			err = registerArandomMessage(stream,100)
			if err != nil {
				t.Errorf("Should not fail in stream send or recive... %+v", err)
			}
		}
	}

}

