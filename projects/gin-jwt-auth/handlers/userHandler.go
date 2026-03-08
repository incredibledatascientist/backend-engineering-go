package handlers

import (
	"fmt"
	"gin-jwt-auth/database"
	"gin-jwt-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// func HashPassword(password string) {

// }
// func VerifyPassword()

func UserSignup(c *gin.Context) {
	req := models.UserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Errorf("invalid requests")})
		return
	}

	db := database.GetDB()

	// Create user
	user := models.User{
		Username: req.Username,
		Password: string(hash),
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	fmt.Println("user:", user)
	c.JSON(http.StatusOK, gin.H{"message": "user created", "user": user})
}

// func UserLogin(c *gin.Context) {

// }

// func GetUsers(c *gin.Context) {

// }

// func GetUser(c *gin.Context) {

// }
