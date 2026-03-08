package routes

import (
	"gin-jwt-auth/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	// routes.Use(middleware.Authenticate())
	routes.GET("/users", handlers.GetUsers)
	routes.GET("/users/:id", handlers.GetUser)
}
