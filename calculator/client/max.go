package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func doMax(client pb.CalculatorServiceClient) {
	log.Println("doMax() was invoked.")

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating request stream: %v\n", err)
	}

	// create the data for the request stream
	numbers := []int32{1, 5, 3, 6, 2, 20}

	waitc := make(chan struct{})

	// goroutine to send the request stream
	go func() {
		for _, number := range numbers {
			log.Printf("Sending request stream: %v\n", number)
			stream.Send(&pb.MaxRequest{Number: number})
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

			log.Printf("Current max number: %v\n", res.MaxNumber)
		}

		close(waitc)
	}()

	<-waitc
}
