package server

import (
	"context"
	"log"
	"net/http"

	"jwt-auth/internal/config"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(cfg config.Config) *HTTPServer {
	// r := chi.NewRouter()

	// register routes
	// RegisterRoutes(r)

	srv := &http.Server{
		Addr: cfg.Server.Addr,
		// Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	return &HTTPServer{server: srv}
}

func (s *HTTPServer) Start() error {
	log.Printf("HTTP server started on addr %s\n", s.server.Addr)

	err := s.server.ListenAndServe()
	if err != nil {
		log.Printf("%v\n", err)
	}

	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return s.server.Shutdown(ctx)
}
