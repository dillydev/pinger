package rest

import (
	"context"
	"net/http"

	pb "github.com/dgparker/pinger/api/protobuf"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Run starts the REST server
func Run() error {
	mux := runtime.NewServeMux()
	err := pb.RegisterPingerHandlerFromEndpoint(context.Background(), mux, "localhost:9000", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}

	return http.ListenAndServe(":9001", mux)
}
