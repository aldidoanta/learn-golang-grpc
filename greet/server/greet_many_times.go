package main

import (
	"fmt"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("greet many times func was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Stream #%d: Hello %s", i, in.FirstName)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
