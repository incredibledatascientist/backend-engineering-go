package routes

import (
	"gin-jwt-auth/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	authGroup := routes.Group("/users")
	{
		authGroup.POST("/signup", handlers.UserSignup)
		authGroup.POST("/login", handlers.UserLogin)
	}
}
