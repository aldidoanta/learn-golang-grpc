package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
)

func createBlog(client pb.BlogServiceClient) string {
	log.Println("createBlog() was invoked.")

	blog := &pb.Blog{
		AuthorId: "Budi",
		Title:    "First blog",
		Content:  "Test content",
	}

	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("A blog item has been created with id: %s\n", res.Id)
	return res.Id
}
