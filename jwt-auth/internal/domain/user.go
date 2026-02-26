package models

import "time"

type UserRole int

const (
	_ = iota
	USER
	ADMIN
	GUEST
)

type User struct {
	ID           int       `json:"id"`
	FirstName    *string   `json:"first_name" validate:"required, min=2, max=100"`
	LastName     *string   `json:"last_name"`
	Username     *string   `json:"username" validate:"required, min=3, max=100"`
	Password     *string   `json:"-" validate:"required, min=5, max=100"`
	Email        *string   `json:"email" validate:"email, required, min=5, max=100"`
	Phone        *string   `json:"phone" validate:"required, min=10"`
	Role         *UserRole `json:"role" validate:"required, min=10"`
	AccessToken  *string   `json:"access_token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
