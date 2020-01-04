package main


import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
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
	//r, err := c.Register(ctx,req)
	var header, trailer metadata.MD
	r, err := c.Register(
		ctx,
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s, %s", r.GetStatus(),r.Id)
}

func registerArandomUser(rc RegisterClient) (Status,error) {
	clientDeadline := time.Now().Add(time.Duration(20) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	req := randomUserForTest(10)
	var header, trailer metadata.MD
	defer cancel()
	r, err := rc.Register(
		ctx,
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		fmt.Printf("Error is %s",err)
		return Status_ERROR,err
	}
	return r.GetStatus(),nil
}

func TestRegisterMaxNumberOfUsers(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)
	cfg :=  createLoginServerConfig()

	for i := 0; i < cfg.MAX_NUMBER_OF_USERS + 100 ; i++ {
		status,err := registerArandomUser(c)
		if err != nil {
			t.Errorf("Should not fail rpc... %+v", err)
		} else if !(status == Status_SUCCESS ||  status == Status_MAXIMUN_NUMBER_OF_USERS_REACHED) {
			t.Errorf("Should eventually fail with Status_MAXIMUN_NUMBER_OF_USERS_REACHED ... %s", status.String())
		}
	}

}






