package main

import (
	pb "github.com/Jennyads/gRPC_Master_Class/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Failed to load CA certificate: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreeetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)

}
