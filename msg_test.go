package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	. "maestro/api"
	"sync"
	"testing"
	"time"

	//	"time"
	"context"
)


func registerArandomMessage(stream Message_MsgClient,cfg  *ServerConfig ) error  {
	numberOfMessages := cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC -	1

	var messagePipe = make(chan *MsgReq,numberOfMessages)
	var sendFail,reciveFail chan *error
	sendFail = make(chan *error,numberOfMessages)
	reciveFail = make(chan *error,numberOfMessages)


	close(messagePipe)


	var wg sync.WaitGroup
	wg.Add(1)

	go func( )  {
		defer wg.Done()
		for i := 0; i < numberOfMessages; i++ {
			r := randomMessageForTest(100,i)
			err := stream.Send(r)
			if err != nil {
				sendFail <- &err
				break
			}
		}
		close(sendFail)
	}()

	wg.Add(1)
	go func()  {
		defer wg.Done()
		for i := 0; i < numberOfMessages; i++ {
			resp,err := stream.Recv()
			fmt.Printf("recived %s\n",resp.Status.String())
			if err != nil || resp.Status != Status_SUCCESS{
				reciveFail <- &err
				break
			}
		}
		close(reciveFail)
	}()




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

	wg.Wait()
	return nil

}


func Test_Msg(t *testing.T) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewMessageClient(conn)
	cfg := createServerConfig()
	cfg.MAX_NUMBER_OF_TOPICS = 20
	cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC = 1000

	for i := 0; i < cfg.MAX_NUMBER_OF_TOPICS  -1 ; i++ {
		clientDeadline := time.Now().Add(time.Duration(30) * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
		defer cancel()
		stream,err := c.Msg(ctx)
		if err != nil {
			t.Errorf("Should not fail rpc... %+v", err)
		}  else {
			err = registerArandomMessage(stream,cfg)
			if err != nil {
				t.Errorf("Should not fail in stream send or recive... %+v", err)
			}
		}
	}

}

