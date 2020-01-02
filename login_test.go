package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
	. "maestro/api"
	"context"
)

func TestLoginFail(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	r := NewRegisterClient(conn)
	c := NewLoginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := randomUserForTest(10)
	resp, err := r.Register(ctx,req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s, %s", resp.GetStatus(),resp.Id)
	}
	ctx, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReq := LoginReq{Id: "",UserName: "usertofail",PassWord: []byte("passwordtofail"),Device:"devicetofail"}
	lresp,err := c.Authenticate(ctx,&loginReq)
	if err != nil {
		t.Errorf("could not authenticate %+v",err)
	}

	if lresp.Status != Status_FAIL {
		t.Errorf("test should fail %s",lresp.Status)
	}
}

func TestLoginSuccess(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	r := NewRegisterClient(conn)
	c := NewLoginClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := randomUserForTest(10)
	ans , err := r.Register(ctx,req)
	if err != nil {
		t.Errorf("could not register user: %v", req)
	} else {
		log.Printf("Registered user %+v\n",ans.Key)
	}
	ctx, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReq := &LoginReq{Id: "",UserName: req.UserName,PassWord: req.PassWord,Device: req.Device}
	lresp,err := c.Authenticate(ctx,loginReq)
	if err != nil {
		t.Errorf("could not authenticate %+v",err)
	}

	fmt.Printf("resp is %+v\n",lresp)

	if lresp.Status != Status_SUCCESS {
		fmt.Printf("failed for user %+v\n",req)
		t.Errorf("test should succeed %+v",lresp)
	}
}

func TestLoginBlock(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	r := NewRegisterClient(conn)
	c := NewLoginClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := randomUserForTest(10)
	_ , err = r.Register(ctx,req)
	if err != nil {
		t.Errorf("could not register user: %v", req)
	} else {
		log.Printf("Registered user %+v\n",req)
	}
	ctx, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReqCorrectPassword := &LoginReq{Id: "",UserName: req.UserName,PassWord: req.PassWord,Device: req.Device}
	loginReqWrongPassword := &LoginReq{Id: "",UserName: req.UserName,PassWord: []byte("wrongPassword"),Device: req.Device}
	lresp,err := c.Authenticate(ctx,loginReqWrongPassword)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}
	lresp,err = c.Authenticate(ctx,loginReqWrongPassword)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}
	lresp,err = c.Authenticate(ctx,loginReqWrongPassword)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}
	lresp,err = c.Authenticate(ctx,loginReqCorrectPassword)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}


	if lresp.Status != Status_BLOCKED {
		fmt.Printf("fail at %+v",lresp)
		t.Errorf("test should block after tree trails with wrong password... %s",lresp.Status)
	}
}





