package main

import (
	"context"
	"fmt"

	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EntityServiceServer) CreateEntity(ctx context.Context, req *pb.CreateEntityReq) (*pb.CreateEntityRes, error) {
	// Essentially doing req.Entity to access the struct with a nil check
	entity := req.GetEntity()
	// Now we have to convert this into a EtityItem type to convert into BSON
	data := EntityItem{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		Name:        entity.GetName(),
		Description: entity.GetDescription(),
		URL:         entity.GetUrl(),
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := entitydb.InsertOne(mongoCtx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to entity, first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	oid := result.InsertedID.(primitive.ObjectID)
	// Convert the object id to it's string counterpart
	entity.Id = oid.Hex()
	// return the entity in a CreateEntityRes type
	return &pb.CreateEntityRes{Entity: entity}, nil
}
