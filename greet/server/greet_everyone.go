package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone() was invoked.")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the request stream: %v\n", err)
		}

		res := fmt.Sprintf("Hello, %v!\n", req.FirstName)
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending the response stream: %v\n", err)
		}
	}
}
