package main

import (
	"context"
	"io"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes() was invoked.")

	req := &pb.GreetRequest{
		FirstName: "Budi",
	}

	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error on GreetManyTimes: %v\n", err)
	}

	for { //infinite loop
		msg, err := stream.Recv()

		if err == io.EOF { //if finished
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
