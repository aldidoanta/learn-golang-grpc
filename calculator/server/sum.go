package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	return &pb.SumResponse{
		Sum: in.FirstInt + in.SecondInt,
	}, nil
}
