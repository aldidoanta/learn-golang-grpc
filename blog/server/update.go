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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog() was called with %v\n", in)

	// parse ID from request
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannor parse ID",
		)
	}

	// prepare the new Blog document
	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": data}
	res, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)

	// general error handling
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	// error handling if ID is not found
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with ID %s\n", in.Id),
		)
	}

	return &emptypb.Empty{}, nil
}
