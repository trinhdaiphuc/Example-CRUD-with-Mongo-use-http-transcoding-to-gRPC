package services

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateEntity is a gRPC function to update an entity in MongoDB
func (s *Entities) UpdateEntity(ctx context.Context, req *pb.UpdateEntityReq) (*pb.UpdateEntityRes, error) {
	// Get the Entity data from the request
	Entity := req.GetEntity()

	// Convert the Id string to a MongoDB UUID
	uuid, err := uuid.FromString(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied Entity id to a MongoDB UUID: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"name":        Entity.GetName(),
		"description": Entity.GetDescription(),
		"url":         Entity.GetUrl(),
	}

	// Convert the uuid into an unordered bson document to search by id
	filter := bson.M{"_id": uuid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := s.EntityCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := &models.EntityItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find Entity with supplied ID: %v", err),
		)
	}
	return &pb.UpdateEntityRes{
		Entity: &pb.Entity{
			Id:          decoded.ID.String(),
			Name:        decoded.Name,
			Description: decoded.Description,
			Url:         decoded.URL,
		},
	}, nil
}
