package main

import (
	"context"
	pb "github.com/Jennyads/gRPC_Master_Class/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Jenny",
		Title:    "My first blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("error while updating: %v\n", err)
	}
	log.Println("Blog was updated successfully")
}
