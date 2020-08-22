package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var conn *grpc.ClientConn

	err := godotenv.Load()

	conn, err = grpc.Dial(os.Getenv("SERVER_HOST"), grpc.WithInsecure())
	log.Println("Host: ", os.Getenv("SERVER_HOST"))
	if err != nil {
		log.Fatalf("Did not connect %s ", err)
	}
	defer conn.Close()

	client := pb.NewEntityServiceClient(conn)

	response, err := client.ListEntities(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error when calling ListEntities: %s", err)
	}

	for {
		entities, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListEntities(_) = _, %v", client, err)
		}
		log.Println(entities)
	}
}
