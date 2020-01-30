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

func ainit() {
	go func() {
		server := NewServer()
		server.Start(true)
		fmt.Printf("Exiting....")
	}()
}


func TestRegisterMaxNumberOfUsers(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)
	//cfg := createServerConfig()



	req := randomUsersForTests(1000, 1000)

	var wg sync.WaitGroup
	for sysIndex, system := range req {
		wg.Add(1)
		go func(sysIndex int,system []*RegisterReq, t *testing.T) {
			defer wg.Done()
			for userIndex, user := range system {
				ctx := context.Background()
				fmt.Printf("Processing for System[%d],User[%d]\n", sysIndex, userIndex)
				_, err := c.Register(
					ctx,
					user, )
				if err != nil {
					t.Errorf("Should not fail %s", err.Error())
				}
			}
		}(sysIndex,system,t)
	}
	wg.Wait()
}
