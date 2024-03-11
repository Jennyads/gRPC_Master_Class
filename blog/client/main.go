package main

import (
	pb "github.com/Jennyads/gRPC_Master_Class/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id) //valid
	//readBlog(c, "aNonExistentID") //invalid
	updateBlog(c, id)
	listBlog(c)
	//deleteBlog(c, id)

}
