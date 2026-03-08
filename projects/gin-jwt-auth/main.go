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

func initDatabase() (*gorm.DB, error) {
	postgres := database.PostgresConfig{
		Name:     "ginjwtauth",
		Port:     5432,
		Host:     "localhost",
		User:     "postgres",
		Password: "infierms",
		TimeZone: "Asia/Kolkata",
		SSLMode:  "disable",
	}

	db, err := database.NewPostgresDB(database.Config{
		Postgres: postgres,
	})
	
	fmt.Println("db:", db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	// Initialize database
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// Run DB migrations
	if db != nil {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}
	fmt.Println("db:", db)

	router := gin.New()
	router.Use(gin.Logger())

	// routes.

	router.GET("/api/v1/health", handlers.HealthHandler)

	fmt.Println("Server is running on addr: localhost:8080")
	router.Run("localhost:8080")

	// Close database connection
	if db != nil {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}

	log.Println("server stopped gracefully")
}
