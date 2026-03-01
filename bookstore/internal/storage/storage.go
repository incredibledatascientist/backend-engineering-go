package storage

import "bookstore/internal/domain"

// type UserStorage interface {
// 	CreateUser(*domain.User) error
// 	UpdateUser(*domain.User) error
// 	GetUser(id int) (*domain.User, error)
// 	GetUsers() ([]domain.User, error)
// }

/*
Improvements need:
-----------------
	1. Whenever we add any new methods we need to add in interface.

// */

// func NewStorage(cfg config.Config) (MovieStorage, error) {
// 	switch cfg.Storage {

// 	case "postgres":
// 		return postgres.NewPostgresStore()

// 	// case "sqlite":
// 	// 	return sqlite.NewSQLiteStore(cfg)

// 	// case "memory":
// 	// 	return memory.NewMemoryStore(), nil

// 	default:
// 		return nil, fmt.Errorf("unsupported storage driver: %s", cfg.Storage)
// 	}
// }

// BookStorage Interface
type BookStorage interface {
	CreateBook(book *domain.Book) error
	GetBooks() ([]domain.Book, error)
	GetBook(id uint) (*domain.Book, error)
	DeleteBook(id uint) error
}
