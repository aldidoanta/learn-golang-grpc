package main

import (
	"context"
	"io"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(client pb.BlogServiceClient) {
	log.Println("listBlog() was invoked.")

	stream, err := client.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listBlog(): %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF { // end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving listBlog() stream: %v\n", err)
		}

		log.Println(res)
	}
}
