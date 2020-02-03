package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	. "maestro/api"
	"sync"
	"testing"
	"time"

	//	"time"
	"context"
)

func Test_Msg(t *testing.T)  {

	postfix := 10000
	token, appName, err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v", err)
	} else {
		t.Logf("token[%s] app[%s]", token, appName)
	}

	cfg := createAppConfig(createServerConfig(), appName)
	numberOfMessages := cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC

	var sendFail, reciveFail chan *error
	sendFail = make(chan *error, numberOfMessages)
	reciveFail = make(chan *error, numberOfMessages)

	clientDeadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", token, "app", appName)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer conn.Close()
	c := NewMsgClient(conn)

	msgs := randomMessageForTest(cfg.MAX_NUMBER_OF_MESSAGES_PER_TOPIC, cfg.MAX_NUMBER_OF_TOPICS)

	/*
	for _,m := range msgs {
		PrettyPrint(m)
	}

	 */

	t.Log("Before stream")

	stream, err := c.Put(ctx)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Log("After stream")

	var wg sync.WaitGroup
	wg.Add(1)

	t.Log("Before go functions")

	go func() {
		defer wg.Done()
		t.Log("In go function for send")
		for _, m := range msgs {
			err := stream.Send(m)
			if err != nil {
				sendFail <- &err
				break
			}
			t.Logf("sending Message[%s]",m.Uuid)
		}
		close(sendFail)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Log("In go function for Rec")
	loop:
		for {
			q, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break loop
				}
				reciveFail <- &err
			} else {
				t.Logf("Recieved Status[%s] for Message[%s]", q.Status.String(), q.Uuid)
			}
		}
		close(reciveFail)
	}()

	t.Log("Waiting for go functions to terminate")

	wg.Wait()

	for e := range sendFail {
		if e != nil {
			t.Logf(err.Error())
			t.Fail()
		}
	}

	for e := range reciveFail {
		if e != nil {
			t.Logf(err.Error())
			t.Fail()
		}
	}

	t.Logf("Success")

}
