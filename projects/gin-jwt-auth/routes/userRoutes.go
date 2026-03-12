package routes

import (
	"gin-jwt-auth/handlers"
	"gin-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	userGroup := routes.Group("/users")
	userGroup.Use(middleware.Authenticate)
	{
		userGroup.GET("", handlers.GetUsers)
		// userGroup.GET("/:id", handlers.GetUser)
	}
}
