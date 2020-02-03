package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	. "maestro/api"
	"sync"
	"testing"
	"time"

	//	"time"
	"context"
)

func registerArandomMessage(cfg *ServerConfig,token,appName string) error {
	numberOfMessages := cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC

	var sendFail, reciveFail chan *error
	sendFail = make(chan *error, numberOfMessages)
	reciveFail = make(chan *error, numberOfMessages)

	clientDeadline := time.Now().Add(30 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,"bearer-bin",token,"app",appName)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewMsgClient(conn)

	stream, err := c.Put(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done()
		for i := 0; i < numberOfMessages; i++ {
			r := randomMessageForTest(100, i)
			err := stream.Send(r)
			if err != nil {
				sendFail <- &err
				break
			}
		}
		close(sendFail)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		loop:
		for {
			_, err = stream.Recv()
			if err != nil {
				if err == io.EOF {
					break loop
					reciveFail <- &err
				}
			}
		}
		close(reciveFail)
	}()

	wg.Wait()

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
	cfg := createServerConfig()
	cfg.MAX_NUMBER_OF_TOPICS = 10
	cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC = 100

	postfix := 10000
	token,app,err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v", err)
	} else {
		t.Logf("token[%s] app[%s]",token,app)
	}

	for i := 0; i < cfg.MAX_NUMBER_OF_TOPICS; i++ {
		err := registerArandomMessage(cfg,token,app)
		if err != nil {
			t.Errorf("Should not fail in stream send or recive... %+v", err)
		}
	}
}
