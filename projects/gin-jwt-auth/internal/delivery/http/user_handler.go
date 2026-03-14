package http

import (
	"errors"
	"net/http"

	"gin-jwt-auth/internal/domain"
	"gin-jwt-auth/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userUsecase domain.UserUsecase
	validate    *validator.Validate
}

// NewUserHandler instantiates the HTTP controllers for User
func NewUserHandler(r *gin.Engine, us domain.UserUsecase, authMiddleware gin.HandlerFunc) {
	handler := &userHandler{
		userUsecase: us,
		validate:    validator.New(),
	}

	authGroup := r.Group("/users")
	{
		authGroup.POST("/signup", handler.Signup)
		authGroup.POST("/login", handler.Login)
	}

	protectedGroup := r.Group("/users")
	protectedGroup.Use(authMiddleware)
	{
		protectedGroup.GET("", handler.GetUsers)
	}
}

func (h *userHandler) Signup(c *gin.Context) {
	var req domain.UserSignupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userUsecase.Signup(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrUsernameExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *userHandler) Login(c *gin.Context) {
	var req domain.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userUsecase.Login(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidLogin) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Optional: set HTTP only cookie here dynamically if needed
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", res.AccessToken, 15*60, "", "", false, true)

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userUsecase.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
