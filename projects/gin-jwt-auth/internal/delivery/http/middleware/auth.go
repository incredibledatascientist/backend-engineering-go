package middleware

import (
	"net/http"
	"strings"

	"gin-jwt-auth/internal/domain"
	"gin-jwt-auth/pkg/token"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates a gin middleware for authenticating requests via JWT
func AuthMiddleware(tokenMaker token.TokenMaker, userUsecase domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or cookie required"})
			return
		}

		claims, err := tokenMaker.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Store user metadata in context
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("user_id", claims.Subject)

		// Could optionally fetch user via usecase if strictly required here,
		// though setting identity claims usually suffices.
		
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	// Try cookie
	tokenStr, err := c.Cookie("Authorization")
	if err == nil && tokenStr != "" {
		return tokenStr
	}

	// Try Header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}

	return ""
}
