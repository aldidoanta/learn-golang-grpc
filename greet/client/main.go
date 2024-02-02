package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/aldidoanta/learn-golang-grpc/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// init `opts` variable to add TLS option
	opts := []grpc.DialOption{}
	tls := true

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	// doGreetManyTimes(client)
	// doLongGreet(client)
	// doGreetEveryone(client)
	// doGreetWithDeadline(client, 1*time.Second)
}
