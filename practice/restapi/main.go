package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("----------- main start -----------")

	connStr := "postgres://postgres:infierms@localhost:5432/restapi?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if db.Ping()

	fmt.Println("----------- main end -------------")
}
