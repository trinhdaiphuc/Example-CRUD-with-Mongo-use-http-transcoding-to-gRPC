package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	entityEndpoint = flag.String("entity_endpoint", "localhost:50051", "endpoint of EntityService")
)

type errorBody struct {
	Message string `json:"message,omitempty"`
	Details struct {
		Error string `json:"-"`
	} `json:"detail"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`
	fmt.Println("Custom HTTP error ->")
	fmt.Println(grpc.ErrorDesc(err), grpc.Code(err), err)

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))
	jErr := json.NewEncoder(w).Encode(errorBody{
		Message: grpc.ErrorDesc(err),
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
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterEntityServiceHandlerFromEndpoint(ctx, mux, *entityEndpoint, dialOpts)
	if err != nil {
		return err
	}

	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunEndPoint(":8080"); err != nil {
		glog.Fatal(err)
	}
}
