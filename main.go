package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/services"
	"google.golang.org/grpc"
)

func main() {
	// Start our listener, 50051 is the default gRPC port
	listener, err := net.Listen("tcp", ":50051")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}
	// Set options, here we can configure things like TLS support
	// Set options, here we can configure things like middleware support
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		))}

	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Create EntityService type
	srv := &services.Entities{}
	// Register the service with the server
	pb.RegisterEntityServiceServer(s, srv)

	db, mongoCtx, err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	// Bind our collection to our global variable for use in other methods
	srv.EntityCollection = models.NewEntityCollection(db)

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
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
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
