package main


import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
	"google.golang.org/grpc"
	. "maestro/api"
)

const (
	address     = "localhost:50051"
)


func init() {
	go func() {
		app := newApp()
		app.start()
		fmt.Printf("Exiting....")
	}()
}



func TestRegisterUser(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := randomUserForTest(10)
	r, err := c.Register(ctx,req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s, %s", r.GetStatus(),r.Id)
}

func TestRegisterMaxNumberOfUsers(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)
	clientDeadline := time.Now().Add(time.Duration(20) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	cfg :=  createLoginServerConfig()
	for i := 0; i < cfg.MAX_NUMBER_OF_USERS; i++ {
		req := randomUserForTest(10)
		r, err := c.Register(ctx, req)
		if err != nil {
			t.Errorf("could not create user %+v",err)
		} else {
			log.Printf("Greeting: %s, %s", r.GetStatus(), r.Id)
		}
	}

	req := randomUserForTest(10)
	r, err := c.Register(ctx, req)
	if err == nil {
		if r.Status != Status_MAXIMUN_NUMBER_OF_USERS_REACHED {
			t.Errorf("should fail on maximum number of users %+v", err)
		}
	}

	<-time.After(60 * time.Second)
}






