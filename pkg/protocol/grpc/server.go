package grpc

import (
	"context"

	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	v1 "github.com/cardenasrjl/emem/pkg/api/v1"
	"github.com/cardenasrjl/emem/pkg/logger"
	"github.com/cardenasrjl/emem/pkg/protocol/grpc/middleware"
)

// RunServer runs gRPC service to publish mem
func RunServer(ctx context.Context, v1API v1.MemServiceServer, v1VAPI v1.VolumeServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterMemServiceServer(server, v1API)
	v1.RegisterVolumeServiceServer(server, v1VAPI)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
