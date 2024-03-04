package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Result of sum: %d\n", res.Result)
}
