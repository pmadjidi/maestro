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
		t.Logf("Should not fail creating a user.. %+v\n", err)
		t.Fail()
	} else {
		t.Logf("token[%s] app[%s]\n", token, appName)
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

	t.Logf("Before stream")

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

	t.Logf("After stream\n")

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
				t.Logf("Send Error [%s]\n", err.Error())
				sendFail <- &err
				return
			}
			t.Logf("Sent Message[%s][%s]\n", m.Uuid, m.Topic)
		}
		t.Logf("Returning from go send function")
	}()

	go func() {
		defer wg.Done()
		defer close(reciveFail)
		t.Logf("In go function for Rec\n")
		for {
			q, err := stream.Recv()

			if err != nil  {
				t.Logf("Recive Error Client [%s]\n", err.Error())
				if err == io.EOF {
					return
				}
				reciveFail <- &err
				return
			} else {
				t.Logf("Recieved Status[%s] for Message[%s]\n", q.Status.String(), q.Uuid)
			}
		}
	}()

	t.Logf("Waiting for go functions to terminate\n")

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



	t.Logf("Success\n")

}


func Test_Timeline(t *testing.T) {

	postfix := 10002
	token, appName, err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v\n", err)
		t.Fail()
	} else {
		t.Logf("token[%s] app[%s]\n", token, appName)
	}


	numberOfMessages := 1000


	var reciveFail chan *error

	reciveFail = make(chan *error, numberOfMessages)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer conn.Close()
	c := NewMsgClient(conn)


	t.Logf("Before stream")

	//ctx, _ := context.WithTimeout(context.Background(),10 * time.Second)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", token, "app", appName,"username","calle1002")
	//defer cancel()

	stream, err := c.TimeLine(ctx,&Empty{})
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer stream.CloseSend()

	t.Logf("After stream\n")

	var wg sync.WaitGroup
	wg.Add(1)

	t.Logf("Before go functions\n")


	go func() {
		defer wg.Done()
		defer close(reciveFail)
		t.Logf("In go function for Rec\n")
		for {
			q, err := stream.Recv()

			if err != nil  {
				if err == io.EOF {
					return
				}
				t.Logf("Recive Error Client [%s]\n", err.Error())
				reciveFail <- &err
				return
			} else {
				t.Logf("Recieved Status[%s] for Message[%s]\n", q.Status.String(), q.Uuid)
			}
		}
	}()

	t.Logf("Waiting for go functions to terminate\n")

	wg.Wait()


	for e := range reciveFail {
		if e != nil {
			t.Logf(err.Error())
			t.Fatal()
		}
	}

	t.Logf("Success\n")

}

