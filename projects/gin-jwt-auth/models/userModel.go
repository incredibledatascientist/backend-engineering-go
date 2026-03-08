package models

import (
	"time"
)

type UserRole int

const (
	RoleUser UserRole = iota + 1
	RoleAdmin
	RoleGuest
)

type User struct {
	// ID        uuid.UUID `json:"id"`
	ID           uint      `json:"id"`
	FirstName    *string   `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string   `json:"last_name,omitempty"`
	Username     *string   `json:"username" validate:"required,min=3,max=100"`
	Password     *string   `json:"-" validate:"required,min=3,max=100"`
	Email        *string   `json:"email" validate:"required,email"`
	Phone        *string   `json:"phone" validate:"required,min=10,max=15"`
	Role         UserRole  `json:"role" validate:"required,oneof=1 2 3"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	AccessToken  *string   `json:"access_token"`
	RefreshToken *string   `json:"refresh_token"`
}
