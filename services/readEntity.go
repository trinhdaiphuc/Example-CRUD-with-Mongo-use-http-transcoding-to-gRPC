package services

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadEntity is a gRPC function to get an entity in MongoDB
func (s *Entities) ReadEntity(ctx context.Context, req *pb.ReadEntityReq) (*pb.ReadEntityRes, error) {
	// convert string id (from proto) to mongoDB UUID
	uuid, err := uuid.FromString(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to UUID: %v", err))
	}

	result := s.EntityCollection.FindOne(ctx, bson.M{"_id": uuid})
	// Create an empty Entity to write our decode result to
	data := &models.EntityItem{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, fmt.Sprintf("Could not find Entity with UUID %s: %v", req.GetId(), err),
		)
	}
	// Cast to ReadEntityRes type
	response := &pb.ReadEntityRes{
		Entity: &pb.Entity{
			Id:          uuid.String(),
			Name:        data.Name,
			Description: data.Description,
			Url:         data.URL,
		},
	}
	return response, nil
}
