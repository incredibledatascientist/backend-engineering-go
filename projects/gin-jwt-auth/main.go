package main

import (
	// "gin-jwt-auth/routes"

	"fmt"
	"gin-jwt-auth/database"
	"gin-jwt-auth/handlers"
	"gin-jwt-auth/models"
	"log"

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

	// router := gin.New()
	// router.Use(gin.Logger())
	router := gin.Default()

	// routes.
	// router.GET("/api/v1/health", handlers.HealthHandler) // API Versioning.
	router.GET("/health", handlers.HealthHandler)

	// User Routes
	router.POST("/users/signup", handlers.UserSignup)

	fmt.Println("Server is running on addr: localhost:8080")
	router.Run("localhost:8080")

	// Close database connection
	if db != nil {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}

	log.Println("server stopped gracefully")
}
