package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/models"
	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/services"
	"google.golang.org/grpc"
)

type errorBody struct {
	Message string `json:"message,omitempty"`
	Details string `json:"detail"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`
	fmt.Println("Message: ", grpc.ErrorDesc(err))
	fmt.Println("Code: ", grpc.Code(err))

	w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))
	jErr := json.NewEncoder(w).Encode(errorBody{
		Message: grpc.ErrorDesc(err),
		Details: strconv.Itoa(int(grpc.Code(err))),
	})

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}

func RunEndPoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println("Starting server on port :8080.")
	mux := runtime.NewServeMux(append(opts, runtime.WithErrorHandler(CustomHTTPError))...)

	srv := &services.Entities{}
	db, mongoCtx, err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	// Bind our collection to our global variable for use in other methods
	srv.EntityCollection = models.NewEntityCollection(db)

	err = pb.RegisterEntityServiceHandlerServer(ctx, mux, srv)
	if err != nil {
		return err
	}

	err = http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println("Error when listen ", err)
	}
	fmt.Println("\nStopping the server...")
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
	return nil
}

func main() {
	defer glog.Flush()

	if err := RunEndPoint(":8080"); err != nil {
		glog.Fatal(err)
	}
}
