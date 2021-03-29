package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/services"
	"github.com/whatvn/denny"
	"github.com/whatvn/denny/middleware/grpc"
	"github.com/whatvn/denny/middleware/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() (*mongo.Client, error) {
	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	mongoURI := "mongodb://localhost:27017"
	if len(os.Getenv("DB_HOST")) > 0 {
		mongoURI = os.Getenv("DB_HOST")
	}

	// non-nil empty context
	mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoURI))

	fmt.Println("DB_HOST ", os.Getenv("DB_HOST"))
	// Handle potential errors
	if err != nil {
		return nil, err
	}

	// Check whether the connection was successful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Mongodb")
	return db, nil
}

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
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	// setup grpc server
	setupGRPCHandler(server, db)

	// then http
	setupHTTPHandler(server, db)

	// start server in dual mode
	if err := server.GraceFulStart(":8080"); err != nil {
		panic(err)
	}
}
