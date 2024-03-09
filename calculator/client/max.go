package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 5, 3, 6, 2, 10}

		for _, number := range numbers {
			log.Printf("Sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
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
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received a new maximum: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc

}
