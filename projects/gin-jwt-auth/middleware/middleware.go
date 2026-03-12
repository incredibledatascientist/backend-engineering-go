package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"gin-jwt-auth/handlers"
	"gin-jwt-auth/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "master@golang" // fallback
	}
	return []byte(secret)
}

func Authenticate(c *gin.Context) {
	fmt.Println("------------- middleware start ----------")
	
	// Attempt to get token from cookie first
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		// Fallback to Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or cookie required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}
		tokenString = parts[1]
	}

	claims := &utils.SignedDetails{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		return
	}

	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Set("user_id", claims.Subject)

	// Optionally fetch full user object if handlers depend on it heavily
	user, errFetch := handlers.GetUser(parseUint(claims.Subject))
	if errFetch == nil && user != nil {
		c.Set("user", user)
	}

	fmt.Println("------------- middleware passed ----------")
	c.Next()
}

// helper
func parseUint(s string) uint {
	var n uint
	fmt.Sscanf(s, "%d", &n)
	return n
}
