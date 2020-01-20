package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	. "maestro/api"
	"sync"
	"testing"
)

const (
	address = "localhost:50051"
)

func init() {
	go func() {
		server := NewServer()
		server.Start(true)
		fmt.Printf("Exiting....")
	}()
}

/*
func TestRegisterUser(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	req := randomUserForTest(1)[0]
	PrettyPrint(req)
	//r, err := c.Register(ctx,req)
	var header, trailer metadata.MD
	_, err = c.Register(
		ctx,
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		log.Fatalf("could not greet: %s", err)
	}
	log.Printf("Greeting: %s, %s",req.FirstName,req.LastName)
}

*/

func TestRegisterMaxNumberOfUsers(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)
	cfg := createServerConfig()

	req := randomUsersForTests(cfg.MAX_NUMBER_OF_USERS, cfg.MAX_NUMBER_OF_APPS)

	var wg sync.WaitGroup
	for sysIndex, system := range req {
		wg.Add(1)
		go func(sysIndex int,system []*RegisterReq, t *testing.T) {
			defer wg.Done()
			for userIndex, user := range system {

				fmt.Printf("Processing for System[%d],User[%d]\n", sysIndex, userIndex)
				_, err := c.Register(
					context.Background(),
					user, )
				if err != nil {
					t.Errorf("Should not fail %s", err.Error())
				}
			}
		}(sysIndex,system,t)
	}
	wg.Wait()
}
