package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	// dsn := "user=postgres dbname=gobank sslmode=verify-full"
	dsn := "user=postgres password=infierms dbname=jwtauth sslmode=disable"
	// dsn := "postgres://postgres:infierms@localhost/gobank?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	store := &PostgresStore{db: db}

	// Create table automatically
	// if err := store.CreateMovieTable(); err != nil {
	// 	return nil, err
	// }

	return store, nil

}
