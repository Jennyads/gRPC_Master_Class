package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes function was invoked")

	req := &pb.PrimeRequest{
		Number: 1239039284,
	}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("Primes: %d\n", res.Result)
	}
}
