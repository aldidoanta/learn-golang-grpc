package main

import (
	"context"
	"io"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func doGetPrimeFactors(client pb.CalculatorServiceClient) {
	log.Printf("doGetPrimeFactors() was invoked.")

	req := &pb.PrimeRequest{
		Number: 300,
	}

	stream, err := client.GetPrimeFactors(context.Background(), req)

	if err != nil {
		log.Fatalf("Error on GetPrimeFactors: %v\n", err)
	}

	for { //infinite loop
		msg, err := stream.Recv()

		if err == io.EOF { //if the stream is finished
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GetPrimeFactors: %d\n", msg.PrimeFactor)
	}
}
