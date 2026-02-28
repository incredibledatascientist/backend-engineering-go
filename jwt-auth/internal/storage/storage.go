package storage

import (
	"fmt"
	"jwt-auth/internal/config"
	"jwt-auth/internal/domain"
	"jwt-auth/internal/storage/postgres"
)

// type UserStorage interface {
// 	CreateUser(*domain.User) error
// 	UpdateUser(*domain.User) error
// 	GetUser(id int) (*domain.User, error)
// 	GetUsers() ([]*domain.User, error)
// }

/*
Improvements need:
-----------------
	1. Whenever we add any new methods we need to add in interface.

*/

type MovieStorage interface {
	CreateMovie(*domain.Movie) error
	// UpdateMovie(*domain.Movie) error
	// GetMovie(id int) (*domain.Movie, error)
	// GetMovies() ([]*domain.Movie, error)
}

func NewStorage(cfg config.Config) (MovieStorage, error) {
	switch cfg.Storage {

	case "postgres":
		return postgres.NewPostgresStore()

	// case "sqlite":
	// 	return sqlite.NewSQLiteStore(cfg)

	// case "memory":
	// 	return memory.NewMemoryStore(), nil

	default:
		return nil, fmt.Errorf("unsupported storage driver: %s", cfg.Storage)
	}
}
