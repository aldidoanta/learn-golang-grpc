package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

var address string = "127.0.0.1:3000"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials((insecure.NewCredentials())))

	if err != nil {
		log.Fatalf("Failed to connect to address %v: %v", address, err)
	}
	defer connection.Close()

	client := pb.NewCalculatorServiceClient(connection)

	// doSum(client)
	doGetPrimeFactors(client)
}
