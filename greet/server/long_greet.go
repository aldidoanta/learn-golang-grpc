package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet() was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF { // end of request stream
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil { // error handling
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving request stream: %v", req.FirstName)

		res += fmt.Sprintf("Hello, %s!\n", req.FirstName) // concatenate the streamed requests
	}
}
