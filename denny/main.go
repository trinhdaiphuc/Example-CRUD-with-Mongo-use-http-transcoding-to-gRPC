package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/services"
	"github.com/whatvn/denny"
	"github.com/whatvn/denny/middleware/grpc"
	"github.com/whatvn/denny/middleware/http"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupGRPCHandler(server *denny.Denny, db *mongo.Client) {
	grpcServer := denny.NewGrpcServer(grpc.ValidatorInterceptor)

	srv := &services.Entities{
		EntityCollection: models.NewEntityCollection(db),
	}
	pb.RegisterEntityServiceServer(grpcServer, srv)
	server.WithGrpcServer(grpcServer)
}

func setupHTTPHandler(server *denny.Denny, db *mongo.Client) {
	entity := server.NewGroup("/")
	entity.BrpcController(&services.Entities{
		EntityCollection: models.NewEntityCollection(db),
	})
}

func main() {
	var server = denny.NewServer(true)
	server.WithMiddleware(gin.Recovery(), http.Logger())
	server.RedirectTrailingSlash = false

	// Initialize MongoDb client
	db, mongoCtx, err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// setup grpc server
	setupGRPCHandler(server, db)

	// then http
	setupHTTPHandler(server, db)

	// start server in dual mode
	if err := server.GraceFulStart(":8082"); err != nil {
		panic(err)
	}

	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
