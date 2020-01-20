package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	. "maestro/api"
	"strconv"
	"testing"
	"time"
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
	_, err = r.Register(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s, %s", req.FirstName, req.LastName)
	}

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	loginReq := &LoginReq{Device: "devicetofail"}
	ctx1 = metadata.AppendToOutgoingContext(ctx1, "username", "usernametofail", "password", "passwordtofail", "app", "Test0")
	_, err = c.Authenticate(
		ctx1,
		loginReq,
	)
	if err != nil {
		t.Errorf("could not authenticate %+v", err)
		t.Fail()
	}
}

func createUser(postfix int, password string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	r := NewRegisterClient(conn)
	ps := strconv.Itoa(postfix)

	req := &RegisterReq{UserName: "kalle" + ps,
		PassWord:  []byte(password),
		FirstName: "Kalle" + ps,
		LastName:  "Svensson" + ps,
		Email:     "kalle" + ps + "@gmail.com",
		Phone:     "08-12 18 " + ps,
		Address:   &RegisterReq_Address{Street: "Tomtebogatan " + ps, City: "Stockholm", State: "Sweden", Zip: "113 " + ps},
		Device:    "device-" + ps,
		AppName:   "Test" + ps,}

	_, err = r.Register(context.Background(), req)

	return err

}

func TestLoginSuccess(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewLoginClient(conn)

	postfix := 1000
	appName := "Test1000"
	//cfg := createServerConfig(appName)
	pass := "theRightPassword"

	err = createUser(postfix, pass)
	if err != nil {
		t.Errorf("should have Status_Fail %s", err.Error())
		t.Fail()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	loginReq := &LoginReq{Device: "device to succeed"}
	ctx = metadata.AppendToOutgoingContext(ctx, "username", "kalle1000", "password", pass, "app", appName)
	var header, trailer metadata.MD // variable to store header and trailer
	_, err = c.Authenticate(
		ctx,
		loginReq,
		grpc.Header(&header),   // will retrieve header
		grpc.Trailer(&trailer), // will retrieve trailer
	)

	if err != nil {
		t.Errorf("could not authenticate %+s", err.Error())
	} else {
		token := header.Get("bearer-bin")
		if len(token) != 0 {
			token := header.Get("bearer-bin")
			fmt.Printf("Token := %s", token)
			//decodeToken(token[0],cfg.SYSTEM_SECRET)
		} else {
			t.Errorf("No token, failing %s", token)
			t.Fail()
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
	c := NewLoginClient(conn)

	postfix := 1001

	err = createUser(postfix, "theRightPassword")
	if err != nil {
		t.Errorf("should have Status_Fail %s", err.Error())
		t.Fail()
	}

	for i := 0; i > cfg.MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		loginReq := &LoginReq{Device: "deviceToFail"}
		ctx = metadata.AppendToOutgoingContext(ctx, "username", "kalle1001", "password", "wrongpassword", "app", "Test1001")
		_, err := c.Authenticate(
			ctx,
			loginReq,
		)

		if i > cfg.MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT && err == nil {
			t.Errorf("could not authenticate %+v", err)
			t.Fail()
		} else {
			t.Logf("Got error %s\n", err.Error())
		}

	}

}

func TestInvalidAppName(t *testing.T) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewLoginClient(conn)


	loginReq := &LoginReq{Device: "deviceToFail"}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "username", "kalle1002", "password", "wrongpassword", "app", "Test1002")
	_, err = c.Authenticate(
		ctx,
		loginReq,
	)
	t.Logf("Got Error %s %s\n", err.Error(),Status_INVALID_APPNAME.String())
	if test:=  err.Error() == Status_INVALID_APPNAME.String(); test {
		t.Fail()
	}
}



