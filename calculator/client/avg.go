package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func doAvg(client pb.CalculatorServiceClient) {
	stream, err := client.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling client.Avg(): %v\n", err)
	}

	for i := 0; i < 5; i++ { // send the request stream
		var number int32 = getRandomNumber()
		req := pb.AvgRequest{Number: number}
		log.Printf("Sending number: %d\n", number)
		stream.Send(&req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv() // close the stream

	if err != nil {
		log.Fatalf("Error while receiving response from Avg(): %v\n", err)
	}

	log.Printf("Response from Avg(): %v\n", res.Average)
}

func getRandomNumber() int32 {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	min := 1
	max := 10
	return int32(randomizer.Intn(max-min+1) + min)
}
