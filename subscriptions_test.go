package main

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	. "maestro/api"
	"sync"
	"testing"
	"time"

	//	"time"
	"context"
	"fmt"
)

func Test_SubscriptionsCreaateTopic(t *testing.T) {

	postfix := 10001
	token, appName, err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v\n", err)
	} else {
		fmt.Printf("token[%s] app[%s]\n", token, appName)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer conn.Close()

	c := NewSubscriptionsClient(conn)

	aTopic := &Topic{
		Id:     uuid.New().String(),
		Tag:    "newTopic",
	}

	ts := append(make([]*Topic, 0), aTopic)


	//ctx, _ := context.WithTimeout(context.Background(),10 * time.Second)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", token)

	resp,err := c.Sub(ctx,&TopicReq{List: ts})

	if err != nil {
		fmt.Printf("Error %s",err.Error())
		t.Fail()

	} else {
		PrettyPrint(resp)
	}

}

func Test_SubscriptionsList(t *testing.T) {

	postfix := 10002
	token, appName, err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v\n", err)
	} else {
		fmt.Printf("token[%s] app[%s]\n", token, appName)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer conn.Close()

	c := NewSubscriptionsClient(conn)

	aTopic := &Topic{
		Id:     uuid.New().String(),
		Tag:    "newTopic",
	}

	aTopic1 := &Topic{
		Id:     uuid.New().String(),
		Tag:    "newTopic1",
	}


	ts := append(make([]*Topic, 0), aTopic,aTopic1)


	//ctx, _ := context.WithTimeout(context.Background(),10 * time.Second)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin", token)

	resp,err := c.Sub(ctx,&TopicReq{List: ts})

	if err != nil {
		fmt.Printf("Error %s",err.Error())
		t.Fail()

		resp,err = c.List(ctx,&Empty{})


	} else {
		PrettyPrint(resp)
	}

}

func Test_SubscriptionsPublishAndList(t *testing.T) {

	postfix := 10003
	token1, appName1, err := createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("Should not fail creating a user.. %+v\n", err)
	} else {
		fmt.Printf("token[%s] app[%s]\n", token1, appName1)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
	defer conn.Close()

	s := NewSubscriptionsClient(conn)
	m := NewMsgClient(conn)


	aTopic := &Topic{
		Id:     uuid.New().String(),
		Tag:    "test1",
	}

	aTopic1 := &Topic{
		Id:     uuid.New().String(),
		Tag:    "test2",
	}

	ts := append(make([]*Topic, 0), aTopic,aTopic1)


	msgs := make([]*MsgReq,0)
	for i := 0; i < 100 ; i++ {
		msg1 := &MsgReq{
			Text:     fmt.Sprintf("Message number [%d]",i),
			Pic:      []byte(RandomString(1000)),
			ParentId:  uuid.New().String(),
			Topic:   "test1"  ,
			TimeName: &timestamp.Timestamp{Seconds: time.Now().Unix(),},
			Status: Status_NEW,
			Uuid:     uuid.New().String(),
		}
		msg2 := &MsgReq{
			Text:     fmt.Sprintf("Message number [%d]",i),
			Pic:      []byte(RandomString(1000)),
			ParentId:  uuid.New().String(),
			Topic:   "test2",
			TimeName: &timestamp.Timestamp{Seconds: time.Now().Unix(),},
			Status: Status_NEW,
			Uuid:     uuid.New().String(),
		}

		msgs = append(msgs,msg1,msg2)
	}




	//ctx, _ := context.WithTimeout(context.Background(),10 * time.Second)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "bearer-bin",token1,"username","kalle10003","app",appName1)
	PrettyPrint(ctx)

	resp,err := s.Sub(ctx,&TopicReq{List: ts})

	if err != nil {
		fmt.Printf("Error %s",err.Error())
		t.Fail()

		resp,err = s.List(ctx,&Empty{})

	} else {
		PrettyPrint(resp)
	}



	var sendFail, reciveFail chan *error
	sendFail = make(chan *error, 100)
	reciveFail = make(chan *error, 100)

	stream, err := m.Put(ctx)
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
				if err == io.EOF {
					return
				}
				fmt.Printf("Recive Error Client [%s]\n", err.Error())
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


	ctx = metadata.AppendToOutgoingContext(context.Background(), "bearer-bin",token1,"username","kalle10003","app",appName1)

	resp,err = s.List(ctx,&Empty{})
	PrettyPrint(resp)


	ctx = metadata.AppendToOutgoingContext(context.Background(), "bearer-bin",token1,"username","kalle10003","app",appName1)
	tlclient,err := m.TimeLine(ctx, &Empty{})

	for {
		q, err := tlclient.Recv()

		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("Recive Error Client [%s]\n", err.Error())
			t.Fail()
			return
		} else {
			fmt.Printf("Recieved Status[%s] for Message[%s]\n", q.Status.String(), q.Uuid)
		}
	}




}







