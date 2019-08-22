package main

import (
	"context"
	"fmt"

	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EntityServiceServer) DeleteEntity(ctx context.Context, req *pb.DeleteEntityReq) (*pb.DeleteEntityRes, error) {
	// Get the ID (string) from the request message and convert it to an Object ID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = entitydb.DeleteOne(ctx, bson.M{"_id": oid})
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete Entity with id %s: %v", req.GetId(), err))
	}
	// Return response with success: true if no error is thrown (and thus document is removed)
	return &pb.DeleteEntityRes{
		Success: true,
	}, nil
}
