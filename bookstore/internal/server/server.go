package server

import (
	"context"
	"log"
	"net/http"

	"bookstore/internal/config"
	"bookstore/internal/storage"
)

type HTTPServer struct {
	server *http.Server
	store  storage.BookStorage
}

func NewHTTPServer(cfg config.Config, store storage.BookStorage) *HTTPServer {
	s := &HTTPServer{
		store: store,
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
