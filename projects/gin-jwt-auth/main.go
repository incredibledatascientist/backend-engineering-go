package main

import (
	// "gin-jwt-auth/routes"

	"fmt"
	"gin-jwt-auth/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	addr := "localhost:8080"
	// gin.Default(addr) // Already logger available with default

	router := gin.New()
	router.Use(gin.Logger())

	// routes.

	router.GET("/api/v1/health", handlers.HealthHandler)

	fmt.Println("Server is running on addr:", addr)
	router.Run(addr)
}
