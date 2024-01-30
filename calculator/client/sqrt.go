package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(client pb.CalculatorServiceClient, number int32) {
	log.Printf("doSqrt() was invoked with %v\n", number)

	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: number,
	})

	if err != nil {
		// check if the error is a gRPC error
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("Invalid argument: the input is possibly a negative number")
				return
			}
		} else {
			log.Fatalf("A non-gRPC error: %v\n", err)
		}

	}

	log.Printf("Square root result: %v\n", res.Result)
}
