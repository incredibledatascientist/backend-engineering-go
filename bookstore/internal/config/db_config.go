package config

import (
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Connect() error {
	// dsn := "user=postgres password=infierms dbname=jwtauth sslmode=disable"

	dsn := "host=localhost user=postgres password=infierms dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	db = d
	return nil
}
