package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var address string = "127.0.0.1:3000"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// init the database client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	// access the database collection
	collection = client.Database("blogdb").Collection("blog")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on address %v.\n", address)
	}

	log.Printf("Listening on address %v.\n", address)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v.\n", err)
	}
}
