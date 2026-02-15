package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(id int) error
	GetAccount(id int) (*Account, error)
	GetAllAccount() ([]*Account, error)
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

// main.go
// // Create database
// err = store.createDatabase()
// if err != nil {
// 	log.Fatal(err)
// }

// func (p *PostgresStore) createDatabase() error {
// 	// dbname=gobank - Not work
// 	// dbname=postgres - Then work and later you can change the dbname to gobank
// 	var exists bool

// 	checkQuery := `
// 		SELECT EXISTS (
// 			SELECT 1 FROM pg_database WHERE datname = 'gobank'
// 		);
// 	`

// 	err := p.db.QueryRow(checkQuery).Scan(&exists)
// 	if err != nil {
// 		return err
// 	}

// 	if !exists {
// 		if _, err := p.db.Exec(`CREATE DATABASE gobank;`); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (p *PostgresStore) createGobankTable() error {
	// CASE SENSITIVE - When use Quotes
	// CASE IN-SENSITIVe - If not use Quotes
	query := `
		CREATE TABLE IF NOT EXISTS gobank (
			id INT PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			number VARCHAR(50),
			balance DECIMAL(10,2));
		`

	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (p *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (p *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (p *PostgresStore) GetAccount(id int) (*Account, error) {
	return nil, nil
}

func (p *PostgresStore) GetAllAccount() ([]*Account, error) {
	return nil, nil
}
