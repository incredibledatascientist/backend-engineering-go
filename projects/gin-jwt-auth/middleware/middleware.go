package middleware

import (
	"fmt"
	"gin-jwt-auth/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = "master@golang"

func Authenticate(c *gin.Context) {
	fmt.Println("------------- middleware start ----------")
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(JWT_SECRET), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		currTime := time.Now().Unix()
		expTime := claims["expiry_time"].(float64)
		userId := claims["user_id"].(float64)
		fmt.Println("curr time:", currTime)
		fmt.Println("User id:", userId, " exp time:", expTime)

		user, err := handlers.GetUser(uint(userId))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println(user.Username)
		fmt.Println(*user)

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

		// if float64(currTime) > expTime {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

	} else {
		fmt.Println(err)
	}

	c.Next()
	fmt.Println("------------- middleware ----------")
}
