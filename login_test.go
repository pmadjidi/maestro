package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReq := &LoginReq{Device:"devicetofail"}
	ctx1 = metadata.AppendToOutgoingContext(ctx1, "username", "usernametofail", "password", "passwordtofail")
	lresp, err := c.Authenticate(
		ctx1,
		loginReq,
	)
	if err != nil {
		t.Errorf("could not authenticate %+v",err)
		t.Fail()
	}

	if lresp != nil && lresp.Status != Status_FAIL {
		t.Errorf("should have Status_Fail %s",lresp.Status)
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
	response , err := r.Register(ctx,req)
	if err != nil {
		t.Errorf("could not register user: %v", req)
	} else {
		log.Printf("Registered user %+v\n",response.Id)
	}
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReq := &LoginReq{Device:"devicetofail"}
	ctx1 = metadata.AppendToOutgoingContext(ctx1, "username",req.UserName, "password", string(req.PassWord))
	lresp, err := c.Authenticate(
		ctx1,
		loginReq,
	)
	if err != nil {
		t.Errorf("could not authenticate %+v",err)
	} else {
		md, val := metadata.FromIncomingContext(ctx1)
		if val {
			token := md.Get("bearer-bin")
			fmt.Printf("Token := %s", token)
			if len(token) == 0 || lresp.Status != Status_SUCCESS {
				t.Fail()
			}
		}
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
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReq := &LoginReq{Device: req.Device}
	ctx1 = metadata.AppendToOutgoingContext(ctx1, "username",req.UserName, "password", "wrongPassword")
	lresp, err := c.Authenticate(
		ctx1,
		loginReq,
	)
	lresp,err = c.Authenticate(ctx1,loginReq)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}
	lresp,err = c.Authenticate(ctx1,loginReq)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}
	lresp,err = c.Authenticate(ctx1,loginReq)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	ctx2 = metadata.AppendToOutgoingContext(ctx2, "username",req.UserName, "password", string(req.PassWord))

	lresp,err = c.Authenticate(ctx2,loginReq)
	if err != nil {
		fmt.Printf("status is: %s",lresp.Status)
		t.Errorf("could not authenticate %+v",err)
	}


	if lresp.Status != Status_BLOCKED {
		fmt.Printf("fail at %+v",lresp)
		t.Errorf("test should block after tree trails with wrong password... %s",lresp.Status)
	}
}





