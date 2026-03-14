package domain

import (
	"context"
	"time"
)

// Role defines the authorization levels
type Role int

const (
	RoleUser Role = iota + 1
	RoleAdmin
	RoleGuest
)

// User represents the core business entity
type User struct {
	ID           uint
	FirstName    string
	LastName     string
	Username     string
	Password     string
	Email        string
	Phone        string
	Role         Role
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// UserSignupReq is the DTO for user registration
type UserSignupReq struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=3"`
}

// UserLoginReq is the DTO for user authentication
type UserLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse is the DTO sent back to clients
type UserResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Role         Role   `json:"role"`
}

// AuthResponse is the DTO for login/signup success
type AuthResponse struct {
	Message      string       `json:"message"`
	User         UserResponse `json:"user,omitempty"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

// UserRepository represents the data storage layer interface
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	UpdateTokens(ctx context.Context, id uint, access string, refresh string) error
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
}

// UserUsecase represents the application business rules interface
type UserUsecase interface {
	Signup(ctx context.Context, req UserSignupReq) (AuthResponse, error)
	Login(ctx context.Context, req UserLoginReq) (AuthResponse, error)
	GetUsers(ctx context.Context) ([]UserResponse, error)
	GetUserByID(ctx context.Context, id uint) (*UserResponse, error)
}
