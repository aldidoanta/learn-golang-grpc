package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(client pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("doGreetWithDeadline() was invoked.")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Budi",
	}
	res, err := client.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("An unexpected gRPC Error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC Error: %v\n", err)
		}
	}

	log.Printf("Greeting: %s\n", res.Result)
}
