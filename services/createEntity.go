package services

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateEntity is a gRPC function to create an entity in MongoDB
func (s *Entities) CreateEntity(ctx context.Context, req *pb.CreateEntityReq) (*pb.CreateEntityRes, error) {
	// Essentially doing req.Entity to access the struct with a nil check
	entity := req.GetEntity()
	// Now we have to convert this into a EtityItem type to convert into BSON
	data := &models.EntityItem{
		ID:          uuid.NewV4(),
		Name:        entity.GetName(),
		Description: entity.GetDescription(),
		URL:         entity.GetUrl(),
	}

	// Insert the data into the database, result contains the newly generated UUID for the new document
	_, err := s.EntityCollection.InsertOne(ctx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// Convert the UUID to it's string counterpart
	entity.Id = data.ID.String()
	// return the entity in a CreateEntityRes type
	return &pb.CreateEntityRes{Entity: entity}, nil
}
