package main

import (
	pb "github.com/Jennyads/gRPC_Master_Class/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	number := in.Number
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number = number / divisor
		} else {
			divisor++
		}
	}
	return nil

}