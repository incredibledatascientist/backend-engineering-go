package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-jwt-auth/config"
	"gin-jwt-auth/internal/domain"
	delivery "gin-jwt-auth/internal/delivery/http"
	"gin-jwt-auth/internal/delivery/http/middleware"
	"gin-jwt-auth/internal/repository/memory"
	"gin-jwt-auth/internal/repository/postgres"
	"gin-jwt-auth/internal/usecase"
	"gin-jwt-auth/pkg/hash"
	"gin-jwt-auth/pkg/token"

	"github.com/gin-gonic/gin"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Setup Repository based on Config (Postgres vs Memory)
	var userRepo domain.UserRepository
	var dbConn *gorm.DB

	if cfg.DBType == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password,
			cfg.Postgres.Name, cfg.Postgres.Port, cfg.Postgres.SSLMode, cfg.Postgres.TimeZone,
		)
		db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
		dbConn = db
		userRepo = postgres.NewUserRepository(db)
		log.Println("Database initialized: PostgreSQL")
	} else {
		userRepo = memory.NewUserRepository()
		log.Println("Database initialized: In-Memory")
	}

	// 3. Setup Utilities
	hasher := hash.NewBcryptHasher(cfg.HashCost)
	tokenMaker := token.NewJWTMaker(cfg.JWTSecret)

	// 4. Setup Usecases
	userUsc := usecase.NewUserUsecase(userRepo, hasher, tokenMaker)

	// 5. Setup Delivery (HTTP/Gin)
	router := gin.Default()
	
	// Default Health Route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "db_type": cfg.DBType})
	})

	authMiddleware := middleware.AuthMiddleware(tokenMaker, userUsc)
	delivery.NewUserHandler(router, userUsc, authMiddleware)

	// 6. Graceful Server Startup
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	go func() {
		log.Printf("Server is running on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Clean up db connection if postgres
	if dbConn != nil {
		sqlDB, _ := dbConn.DB()
		_ = sqlDB.Close()
	}

	log.Println("Server exiting gracefully")
}
