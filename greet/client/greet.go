package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Printf("doGreet() was invoked.")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Budi",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n.", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
