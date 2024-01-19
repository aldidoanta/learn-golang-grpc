package main

import (
	"log"
	"net"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on the connection: %v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
