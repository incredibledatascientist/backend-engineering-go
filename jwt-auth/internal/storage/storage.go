package storage

import "jwt-auth/internal/domain"

type UserStorage interface {
	CreateUser(*domain.User) error
	UpdateUser(*domain.User) error
	GetUser(id int) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
}
