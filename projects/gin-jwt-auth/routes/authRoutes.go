package routes

import (
	"gin-jwt-auth/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("/users/signup", handlers.UserSignup)
	routes.POST("/users/login", handlers.UserLogin)
}
