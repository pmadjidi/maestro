package main

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	. "maestro/api"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)



func TestClient(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewRegisterClient(conn)

	/* Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	 */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &RegisterReq{UserName: "Payam",
		PassWord: []byte("abrakatabra"),
	FirstName: "Payam",
	LastName: "Madjidi",
	Email: "pmadjidi@gmail.com",
	Phone: "0708121806",
	Address: &RegisterReq_Address{Street:"Tomtebogatan",City: "Stockholm",State: "Sweden",Zip: "11338"},
	Device: "payams-mac-0"}
	r, err := c.Register(ctx,req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s, %s", r.GetStatus(),r.Id)
}
