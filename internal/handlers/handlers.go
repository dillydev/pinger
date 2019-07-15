package handlers

import (
	"context"

	protobuf "github.com/dgparker/pinger/api/protobuf"
)

// Server implements the gRPC server interface
type Server struct{}

// New ...
func New() (*Server, error) {
	return &Server{}, nil
}

// Ping ...
func (s *Server) Ping(context.Context, *protobuf.PingRequest) (*protobuf.PingResponse, error) {
	return &protobuf.PingResponse{
		Message: "update two",
	}, nil
}
