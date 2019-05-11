package rpc

import (
	"log"
	"net"

	"github.com/dgparker/pinger/internal/handlers"
	pb "github.com/dgparker/pinger/api/protobuf"

	"google.golang.org/grpc"
)

// Run starts the RPC server
func Run() error {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}
	defer func() {
		err := lis.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	hndlrs, err := handlers.New()
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingerServer(grpcServer, hndlrs)

	return grpcServer.Serve(lis)
}
