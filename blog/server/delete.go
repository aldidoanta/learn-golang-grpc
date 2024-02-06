package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog() was called with %v\n", in)

	// parse ID from request
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	// delete a document, applying a filter by ID
	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(
		ctx,
		filter,
	)

	// general error handling
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot delete object in MongoDB",
		)
	}

	// error handling if ID is not found
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with ID %s\n", in.Id),
		)
	}

	return &emptypb.Empty{}, nil
}
