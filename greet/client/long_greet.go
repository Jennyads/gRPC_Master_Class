package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{
			FirstName: "John",
		},
		{
			FirstName: "Jane",
		},
		{
			FirstName: "Joe",
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}
	log.Printf("LongGreet: %v\n", res)
}
