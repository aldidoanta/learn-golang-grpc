package main

import (
	"io"
	"log"
	"math"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max() was invoked.")
	var max int32 = math.MinInt32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the request stream: %v\n", err)
		}

		// compare the number with the current max
		if req.Number > max {
			max = req.Number
			log.Printf("Current max number: %v\n", max)
			err = stream.Send(&pb.MaxResponse{
				MaxNumber: max,
			})

			if err != nil {
				log.Fatalf("Error while sending the response stream: %v\n", err)
			}
		}
	}
}
