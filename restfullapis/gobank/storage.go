/*
Improvements need:
-----------------
	1. Whenever we add any new methods we need to add in interface.

*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(id int) error
	DeleteAllAccount() error
	GetAccount(id int) (*Account, error)
	GetAllAccount() ([]*Account, error)
	GetAccountByNumber(number string) (*Account, error)
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

func (p *PostgresStore) createAccountTable() error {
	// CASE SENSITIVE - When use Quotes
	// CASE IN-SENSITIVe - If not use Quotes
	query := `
		CREATE TABLE IF NOT EXISTS account (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			password VARCHAR(255),
			balance DECIMAL(10,2),
			number VARCHAR(50),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
		`

	_, err := p.db.Exec(query)
	return err
}

func GenerateAccountNumber(id int64) string {
	return fmt.Sprintf("%010d", id)
}

func (p *PostgresStore) CreateAccount(acc *Account) error {
	query := `
		INSERT INTO account (first_name, last_name, password, balance)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	err := p.db.QueryRow(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Password,
		acc.Balance,
	).Scan(&acc.Id)

	acc.Number = GenerateAccountNumber(int64(acc.Id))

	query = `UPDATE account SET number = $1 WHERE id = $2`
	_, err = p.db.Query(query, acc.Number, acc.Id)

	return err
}

func (p *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (p *PostgresStore) DeleteAccount(id int) error {
	_, err := p.GetAccount(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM account WHERE id=$1`
	_, err = p.db.Query(query, id)
	return err
}

func (p *PostgresStore) DeleteAllAccount() error {
	query := `DELETE FROM account`
	_, err := p.db.Query(query)
	return err
}

func (p *PostgresStore) GetAccount(id int) (*Account, error) {
	query := `SELECT * FROM account WHERE id=$1`
	rows, err := p.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		return account, err
	}
	return nil, fmt.Errorf("Account %d not found", id)
}

func (p *PostgresStore) GetAccountByNumber(number string) (*Account, error) {
	query := `SELECT * FROM account WHERE number=$1`
	rows, err := p.db.Query(query, number)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		return account, err
	}
	return nil, fmt.Errorf("Account by number [%s] not found", number)
}

func (p *PostgresStore) GetAllAccount() ([]*Account, error) {
	// Write query params in same order as table fields
	// Id, FirstName, LastName, Number, Balance, CreatedAt.

	query := `SELECT * FROM account`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var accounts []*Account
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := &Account{}
	err := rows.Scan(
		&account.Id,
		&account.FirstName,
		&account.LastName,
		&account.Password,
		&account.Balance,
		&account.Number,
		&account.CreatedAt,
	)

	return account, err
}
