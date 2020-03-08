package services

import (
	"context"
	"fmt"

	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadEntity is a gRPC function to get an entity in MongoDB
func (s *EntityServiceServer) ReadEntity(ctx context.Context, req *pb.ReadEntityReq) (*pb.ReadEntityRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	result := s.EntityCollection.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty EntityItem to write our decode result to
	data := &models.EntityItem{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Entity with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadEntityRes type
	response := &pb.ReadEntityRes{
		Entity: &pb.Entity{
			Id:          oid.Hex(),
			Name:        data.Name,
			Description: data.Description,
			Url:         data.URL,
		},
	}
	return response, nil
}
