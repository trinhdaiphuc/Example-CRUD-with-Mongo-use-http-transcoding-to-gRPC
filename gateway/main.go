package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type errorBody struct {
	Message string `json:"message,omitempty"`
	Details string `json:"detail"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`
	fmt.Println("Custom HTTP error ->")
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
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	entityEndpoint := flag.String("entity_endpoint", os.Getenv("ENTITY_SERVER_HOST"), "endpoint of EntityService")

	err := pb.RegisterEntityServiceHandlerFromEndpoint(ctx, mux, *entityEndpoint, dialOpts)
	if err != nil {
		return err
	}

	err = http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println("Error when listen ", err)
	}
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunEndPoint(":8080"); err != nil {
		glog.Fatal(err)
	}
}
