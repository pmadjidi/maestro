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
	req := randomUserForTest(1)[0]
	_, err = r.Register(ctx,req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s, %s",req.FirstName,req.LastName)
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
	req := randomUserForTest(1)[0]
	_ , err = r.Register(ctx,req)
	if err != nil {
		t.Errorf("could not register user: %v", req)
	} else {
		log.Printf("Registered user %s %s\n",req.FirstName,req.LastName)
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
	cfg := createServerConfig()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	r := NewRegisterClient(conn)
	c := NewLoginClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := randomUserForTest(1)[0]
	_, err = r.Register(ctx, req)
	if err != nil {
		t.Errorf("could not register user: %v", req)
	} else {
		log.Printf("Registered user %+v\n", req)
	}

	for i := 0; i > cfg.MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		loginReq := &LoginReq{Device: req.Device}
		ctx = metadata.AppendToOutgoingContext(ctx, "username", req.UserName, "password", "wrongPassword")
		lresp, err := c.Authenticate(
			ctx,
			loginReq,
		)

		if err != nil {
			fmt.Printf("status is: %s", lresp.Status)
			t.Errorf("could not authenticate %+v", err)
		}
		if  i > cfg.MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT && lresp.Status != Status_BLOCKED {
			t.Errorf("status shoud be blocked,%s", Status_BLOCKED.String())
		}
	}

}






