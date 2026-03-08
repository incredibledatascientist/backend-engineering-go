package handlers

import (
	"context"
	"fmt"
	"gin-jwt-auth/database"
	"gin-jwt-auth/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// func HashPassword(password string) {

// }
// func VerifyPassword()

var JWT_SECRET = "master@golang"

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.ID,
		"expiry_time": time.Now().Add(5 * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 10*60, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}

func GetUsers(c *gin.Context) {
	fmt.Println("----- get user:")
	db := database.GetDB()
	ctx := context.Background()
	users, err := gorm.G[models.User](db).Find(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(id uint) (*models.User, error) {
	db := database.GetDB()
	ctx := context.Background()
	user, err := gorm.G[models.User](db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

// func GetUser(c *gin.Context) {

// }
