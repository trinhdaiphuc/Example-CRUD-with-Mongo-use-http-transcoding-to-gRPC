package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// Global variables for db connection , collection and context

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	err := godotenv.Load()
	flag.Parse()
	defer glog.Flush()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	// Start our listener, 50051 is the default gRPC port
	listener, err := net.Listen("tcp", ":50051")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}
	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Create EntityService type
	srv := &services.EntityServiceServer{}
	// Register the service with the server
	pb.RegisterEntityServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
	mongoctx := context.Background()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err := mongo.Connect(mongoctx, options.Client().ApplyURI(os.Getenv("DB_HOST")))
	fmt.Println("DB_HOST ", os.Getenv("DB_HOST"))
	// Handle potential errors
	if err != nil {
		log.Fatal(err)
		glog.Fatal(err)
	}

	// Check whether the connection was succesful by pinging the MongoDB server
	err = db.Ping(mongoctx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	// Bind our collection to our global variable for use in other methods
	srv.EntityCollection = models.NewEntityCollection(db)

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
			glog.Fatal(err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoctx)
	fmt.Println("Done.")
}
