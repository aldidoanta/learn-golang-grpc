package main

import (
	"io"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg() was invoked.")

	var sum int32 = 0
	var count int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF { // end of request stream
			avg := float64(sum) / float64(count)
			log.Printf("Sent response: %v\n", avg)
			return stream.SendAndClose(&pb.AvgResponse{
				Average: avg,
			})
		}

		if err != nil { // handles error
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)

		// stores the sum and the number of integers from the client stream
		sum += req.Number
		count++
	}
}
