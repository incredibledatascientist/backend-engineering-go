package server

import (
	"context"
	"log"
	"net/http"

	"jwt-auth/internal/config"
	"jwt-auth/internal/handler"
)

type HTTPServer struct {
	server       *http.Server
	movieHandler *handler.MovieHandler
}

func NewHTTPServer(cfg config.Config, movieHandler *handler.MovieHandler) *HTTPServer {
	s := &HTTPServer{
		movieHandler: movieHandler,
	}

	s.server = &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      s.Routes(),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	return s
}

func (s *HTTPServer) Start() error {
	log.Printf("HTTP server started on addr %s\n", s.server.Addr)

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("Unexpected server error: %v\n", err)
		return err
	}

	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return s.server.Shutdown(ctx)
}
