package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {

	log.Println("GreetEveryone was invoked")
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v\n", err)
	}

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

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)

		}
		close(waitc)
	}()
	<-waitc
}
