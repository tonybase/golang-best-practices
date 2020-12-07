package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server is a grpc server wrapped.
type Server struct {
	srv *grpc.Server
}

// NewServer new a gRPC server..
func NewServer() *Server {
	return &Server{srv: grpc.NewServer()}
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {
	log.Println("[GRPC] Listening on: :9000")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}
	go func() {
		if err := s.srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	log.Println("[GRPC] Stopping")

	s.srv.GracefulStop()
	return nil
}
