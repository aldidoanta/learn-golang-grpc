package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked.")

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	// create the data for the request stream
	reqs := []*pb.GreetRequest{
		{FirstName: "Budi"},
		{FirstName: "Tejo"},
		{FirstName: "Sukimin"},
	}

	waitc := make(chan struct{})

	// goroutine to send the request stream
	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// goroutine to receive the response stream
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving response stream: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
