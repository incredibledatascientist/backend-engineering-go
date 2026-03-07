package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bookstore/internal/config"
	"bookstore/internal/domain"
	"bookstore/internal/server"
	"bookstore/internal/storage"
	"bookstore/internal/storage/postgres"

	"gorm.io/gorm"
)

func initStorage(cfg config.Config) (storage.BookStorage, *gorm.DB, error) {
	switch cfg.Storage {
	case config.StoragePostgres:
		// Initialize Postgres database
		db, err := postgres.NewPostgresDB(cfg)
		if err != nil {
			return nil, nil, err
		}
		return postgres.NewBookStore(db), db, nil

	// case config.StorageSQLite:
	// 	// Initialize SQLite database
	// 	db, err := sqlite.NewSQLiteDB(cfg)
	// 	if err != nil {
	// 		return nil, nil, err
	// 	}
	// 	return sqlite.NewBookStore(db), db, nil

	default:
		return nil, nil, fmt.Errorf("unsupported storage type: %s", cfg.Storage)
	}
}

func main() {

	// Parse config file flag
	configFile := flag.String("config", "configs/local.yaml", "Configuration file")
	flag.Parse()

	// Load application configuration
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize storage based on config
	store, db, err := initStorage(cfg)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	// Run DB migrations
	if db != nil {
		if err := db.AutoMigrate(&domain.Book{}); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}

	// Create HTTP server instance
	httpServer := server.NewHTTPServer(cfg, store)

	// Start HTTP server in background
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Printf("HTTP server error: %v\n", err)
		}
	}()

	// Listen for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown HTTP server gracefully
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %v\n", err)
	}

	// Close database connection
	if db != nil {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}

	log.Println("server stopped gracefully")
}
