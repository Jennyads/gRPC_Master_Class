package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
