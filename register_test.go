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
		fmt.Printf("***Exiting....")
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



