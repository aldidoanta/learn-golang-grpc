package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Printf("doSum() was invoked.")

	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstInt:  2,
		SecondInt: 3,
	})

	if err != nil {
		log.Fatalf("Sum failed: %v.\n", err)
	}

	log.Printf("Response of Sum endpoint: %d\n.", res.Sum)
}
