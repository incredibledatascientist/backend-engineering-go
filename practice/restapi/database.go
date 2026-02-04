package database

import (
	"database/sql"
	"fmt"
	"log"
)

var Hostname = "localhost"
var Port = 5432
var Username = "postgres"
var Password = "infierms"
var Database = "restapi" // Database name

func ConnectPostgres() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)
	// All of Hostname, Port, Username, Password, and Database are global variables defined
	// elsewhere in the package—they contain the connection details.
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}
