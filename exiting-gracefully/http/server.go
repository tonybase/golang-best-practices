package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
)

// Server is a HTTP server.
type Server struct {
	srv *http.Server
}

// NewServer new a HTTP server.
func NewServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "okay")
		},
	))
	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	return &Server{srv: srv}
}

// Start start the HTTP server.
func (s *Server) Start(context.Context) error {
	log.Printf("[HTTP] Listening on: %s\n", s.srv.Addr)

	ln, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	go func() {
		if err := s.srv.Serve(ln); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	return nil
}

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	log.Printf("[HTTP] Stopping\n")

	return s.srv.Shutdown(ctx)
}
