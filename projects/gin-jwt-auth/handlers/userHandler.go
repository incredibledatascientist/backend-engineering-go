package handlers

import (
	"context"
	"fmt"
	"gin-jwt-auth/database"
	"gin-jwt-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func UserLogin(c *gin.Context) {
	req := models.UserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	ctx := context.Background()
	user, err := gorm.G[models.User](db).Where("username = ?", req.Username).First(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password not matched"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// func GetUsers(c *gin.Context) {

// }

// func GetUser(c *gin.Context) {

// }
