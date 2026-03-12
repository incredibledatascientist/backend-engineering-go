package handlers

import (
	"net/http"
	"strconv"

	"gin-jwt-auth/database"
	"gin-jwt-auth/models"
	"gin-jwt-auth/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func UserSignup(c *gin.Context) {
	req := models.UserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	db := database.GetDB()

	// Check if user exists
	var count int64
	db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Password: string(hash),
		Role:     models.RoleUser, // Assign default role
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userIdStr := strconv.Itoa(int(user.ID))
	accessToken, refreshToken := utils.GenerateAuthTokens(userIdStr, user.Username, strconv.Itoa(int(user.Role)))

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	db.Save(&user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func UserLogin(c *gin.Context) {
	req := models.UserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	var user models.User
	err = db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	isValid, _ := utils.VerifyPassword(user.Password, req.Password)
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	userIdStr := strconv.Itoa(int(user.ID))
	accessToken, refreshToken := utils.GenerateAuthTokens(userIdStr, user.Username, strconv.Itoa(int(user.Role)))

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	db.Save(&user) // update tokens in db

	// Set cookies for security alternative
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", accessToken, 15*60, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func GetUsers(c *gin.Context) {
	db := database.GetDB()
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Filter out sensitive data here in a real app or use a response struct
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(id uint) (*models.User, error) {
	db := database.GetDB()
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
