package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	. "maestro/api"
	"sync"
	"testing"
	//	"time"
	"context"
	"fmt"
)

func Test_Msg(t *testing.T) {

	postfix := 10000
	token, appName, err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v\n", err)
	} else {
		fmt.Printf("token[%s] app[%s]\n", token, appName)
	}

	cfg := createAppConfig(createServerConfig(), appName)
	numberOfMessages := cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC

	var sendFail, reciveFail chan *error
	sendFail = make(chan *error, numberOfMessages)
	reciveFail = make(chan *error, numberOfMessages)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer conn.Close()
	c := NewMsgClient(conn)

	msgs := randomMessageForTest(cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC, cfg.MAX_NUMBER_OF_TOPICS)

	fmt.Printf("Before stream")

	//ctx, _ := context.WithTimeout(context.Background(),10 * time.Second)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", token, "app", appName)
	//defer cancel()

	stream, err := c.Put(ctx)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer stream.CloseSend()

	fmt.Printf("After stream\n")

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Printf("Before go functions\n")

	go func() {
		defer wg.Done()
		defer close(sendFail)
		defer stream.CloseSend()
		fmt.Printf("In go function for send\n")
		for _, m := range msgs {
			err := stream.Send(m)
			if err != nil && err != io.EOF {
				fmt.Printf("Send Error [%s]\n", err.Error())
				sendFail <- &err
				return
			}
			fmt.Printf("Sent Message[%s][%s]\n", m.Uuid, m.Topic)
		}
		fmt.Printf("Returning from go send function")
	}()

	go func() {
		defer wg.Done()
		defer close(reciveFail)
		fmt.Printf("In go function for Rec\n")
		for {
			q, err := stream.Recv()

			if err != nil  {
				fmt.Printf("Recive Error Client [%s]\n", err.Error())
				if err == io.EOF {
					return
				}
				reciveFail <- &err
				return
			} else {
				fmt.Printf("Recieved Status[%s] for Message[%s]\n", q.Status.String(), q.Uuid)
			}
		}
	}()

	fmt.Printf("Waiting for go functions to terminate\n")

	wg.Wait()


		for e := range sendFail {
			if e != nil {
				t.Logf(err.Error())
				t.Fatal()
			}
		}

		for e := range reciveFail {
			if e != nil {
				t.Logf(err.Error())
				t.Fatal()
			}
		}



	fmt.Printf("Success\n")

}
