package main

import (
	"context"
	"fmt"

	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/server/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EntityServiceServer) UpdateEntity(ctx context.Context, req *pb.UpdateEntityReq) (*pb.UpdateEntityRes, error) {
	// Get the Entity data from the request
	Entity := req.GetEntity()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(Entity.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied Entity id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"name":        Entity.GetName(),
		"description": Entity.GetDescription(),
		"url":         Entity.GetUrl(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := entitydb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := EntityItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find Entity with supplied ID: %v", err),
		)
	}
	return &pb.UpdateEntityRes{
		Entity: &pb.Entity{
			Id:          decoded.ID.Hex(),
			Name:        decoded.Name,
			Description: decoded.Description,
			Url:         decoded.URL,
		},
	}, nil
}
