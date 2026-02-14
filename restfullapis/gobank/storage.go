package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	// UpdateAccount(*Account) error
	// DeleteAccount(id int) error
	// GetAccount(id int) (*Account, error)
	// GetAllAccount() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	// dsn := "user=postgres dbname=gobank sslmode=verify-full"
	dsn := "user=postgres password=infierms dbname=gobank sslmode=disable"
	// dsn := "postgres://postgres:infierms@localhost/gobank?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil

}

func (p *PostgresStore) CreateAccount(*Account) error {
	return nil
}
