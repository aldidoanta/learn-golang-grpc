package main

import (
	"context"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog() was called with %v\n", in)

	// parse ID from request
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannor parse ID",
		)
	}

	// find the Blog document using the ID
	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)

	// unmarshal the document (if any) and check for errors
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with the ID provided",
		)
	}
	return documentToBlog(data), nil
}
