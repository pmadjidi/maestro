package main

import (
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	. "maestro/api"
	"testing"
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

