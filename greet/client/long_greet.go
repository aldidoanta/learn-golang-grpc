package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func doLongGreet(client pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{FirstName: "Budi"},
		{FirstName: "Tejo"},
		{FirstName: "Sukimin"},
	}
	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doLongGreet(): %v\n", err)
	}

	for _, req := range reqs { // send the request stream
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv() // close the stream and receive the server's response

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet(): %v\n", err)
	}

	log.Printf("Response from LongGreet(): %s\n", res.Result)
}
