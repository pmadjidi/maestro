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

func registerArandomMessage(cfg *ServerConfig) error {
	numberOfMessages := cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC

	var sendFail, reciveFail chan *error
	sendFail = make(chan *error, numberOfMessages)
	reciveFail = make(chan *error, numberOfMessages)

	clientDeadline := time.Now().Add(30 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewMessageClient(conn)
	stream, err := c.Msg(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < numberOfMessages; i++ {
			r := randomMessageForTest(100, i)
			err := stream.Send(r)
			if err != nil {
				sendFail <- &err
				break
			}
		}
		close(sendFail)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < numberOfMessages; i++ {
			resp, err := stream.Recv()
			fmt.Printf("recived %s\n", resp.Status.String())
			if err != nil || resp.Status != Status_SUCCESS {
				reciveFail <- &err
				break
			}
		}
		close(reciveFail)
		wg.Done()
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

	fmt.Printf("********\n")
	return nil

}

func Test_Msg(t *testing.T) {
	cfg := createServerConfig()
	cfg.MAX_NUMBER_OF_TOPICS = 10
	cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC = 100

	for i := 0; i < cfg.MAX_NUMBER_OF_TOPICS; i++ {
		err := registerArandomMessage(cfg)
		if err != nil {
			t.Errorf("Should not fail in stream send or recive... %+v", err)
		}
	}
}
