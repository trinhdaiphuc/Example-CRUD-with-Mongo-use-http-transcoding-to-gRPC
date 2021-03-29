package services

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteEntity is a gRPC function to delete an entity in MongoDB
func (s *Entities) DeleteEntity(ctx context.Context, req *pb.DeleteEntityReq) (*pb.DeleteEntityRes, error) {
	// Get the ID (string) from the request message and convert it to an UUID
	uuid, err := uuid.FromString(req.GetId())
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to UUID: %v", err))
	}

	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = s.EntityCollection.DeleteOne(ctx, bson.M{"_id": uuid})
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete Entity with id %s: %v", req.GetId(), err))
	}
	// Return response with success: true if no error is thrown (and thus document is removed)
	return &pb.DeleteEntityRes{
		Success: true,
	}, nil
}
