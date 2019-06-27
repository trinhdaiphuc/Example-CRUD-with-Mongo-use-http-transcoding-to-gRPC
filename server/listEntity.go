package main

import (
	"context"
	"fmt"

	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/server/entity"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EntityServiceServer) ListEntities(req *pb.ListEntitiesReq, stream pb.EntityService_ListEntitiesServer) error {
	// Initiate a EntityItem type to write decoded data to
	data := &EntityItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := entitydb.Find(context.Background(), bson.M{})
	if cursor == nil {
		status.New(codes.FailedPrecondition, "No users have been created")
	}
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	var Entities []*pb.Entity
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send Entity over stream
		entityItem := &pb.Entity{
			Id:          data.ID.Hex(),
			Name:        data.Name,
			Description: data.Description,
			Url:         data.URL,
		}

		Entities = append(Entities, entityItem)
	}

	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	stream.Send(&pb.ListEntitiesRes{
		Entity: Entities,
	})
	return nil
}
