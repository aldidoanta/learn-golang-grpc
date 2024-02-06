package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
)

func readBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog() was invoked.")

	req := &pb.BlogId{Id: id}
	res, err := client.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("A blog was read: %v\n", res)
	return res
}
