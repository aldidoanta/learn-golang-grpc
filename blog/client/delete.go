package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
)

func deleteBlog(client pb.BlogServiceClient, id string) {
	log.Println("deleteBlog() was invoked.")

	_, err := client.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting Blog: %v\n", err)
	}

	log.Println("The blog was deleted")
}
