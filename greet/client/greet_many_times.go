package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/greet/proto"
	"io"
	"log"
)

func doGreeetManyTimes(c pb.GreetServiceClient) {
	log.Printf("GreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "John",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
