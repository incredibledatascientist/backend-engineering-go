package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"jwt-auth/internal/config"
	"jwt-auth/internal/server"
	"jwt-auth/internal/storage"
)

func main() {
	// Parse flags
	configFile := flag.String("config", "configs/local.yaml", "Configuration file")
	flag.Parse()

	// Load config
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create DB store
	_, err = storage.NewStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations & create tables

	// Create server
	httpServer := server.NewHTTPServer(cfg)

	// Start server in goroutine
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Printf("HTTP Server error(%v)\n", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP Server shutdown failed: %v\n", err)
		return
	}

	log.Println("HTTP server stopped gracefully")
}
