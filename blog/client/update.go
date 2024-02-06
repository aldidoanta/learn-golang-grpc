package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
)

func updateBlog(client pb.BlogServiceClient, id string) {
	log.Println("updateBlog() was invoked.")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Charlie",
		Title:    "Edited, Second blog",
		Content:  "Edited Test content",
	}

	_, err := client.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error while updating Blog: %v\n", err)
	}

	log.Println("The blog was updated")
}
