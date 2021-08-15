package services

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListEntities is a gRPC function to list all entities in MongoDB
func (s *Entities) ListEntities(ctx context.Context, req *empty.Empty) (*pb.ListEntitiesRes, error) {
	// Initiate a Entity type to write decoded data to
	data := &models.EntityItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := s.EntityCollection.Find(context.Background(), bson.M{})
	if cursor == nil {
		status.New(codes.FailedPrecondition, "No users have been created")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	res := &pb.ListEntitiesRes{}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send Entity over stream
		res.Entities = append(res.Entities, &pb.Entity{
			Id:          data.ID.String(),
			Name:        data.Name,
			Description: data.Description,
			Url:         data.URL,
		})
	}

	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	return res, nil
}
