package main

import (
	"log"
	"net"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var address string = "127.0.0.1:3000"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on address %v.\n", address)
	}

	log.Printf("Listening on address %v.\n", address)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v.\n", err)
	}
}
