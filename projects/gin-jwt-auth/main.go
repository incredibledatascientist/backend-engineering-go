package main

import (
	"context"
	"fmt"
	"gin-jwt-auth/database"
	"gin-jwt-auth/handlers"
	"gin-jwt-auth/models"
	"gin-jwt-auth/routes" // Ensure this is imported
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Initdb *gorm.DB

func main() {

	// Initialize database
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// Run DB migrations
	if db != nil {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}

	router := gin.Default()

	// API versioning / Health check
	router.GET("/health", handlers.HealthHandler)

	// Register defined routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Configure HTTP server with Graceful Shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run server in a goroutine
	go func() {
		fmt.Println("Server is running on addr: localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 5-second context timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}

	// Close database connection cleanly
	if db != nil {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
			log.Println("Database connection closed gracefully")
		}
	}

	log.Println("Server exiting")
}
