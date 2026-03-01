package sqlite

import (
	"bookstore/internal/config"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB(cfg config.Config) (*gorm.DB, error) {

	dbCfg := cfg.SQLite

	// Validate DB file name
	if dbCfg.Name == "" {
		return nil, fmt.Errorf("[sqlite] database name is required")
	}

	// Open SQLite database file
	db, err := gorm.Open(sqlite.Open(dbCfg.Name), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(1)            // SQLite allows single writer
	sqlDB.SetMaxIdleConns(1)            // Keep one idle connection
	sqlDB.SetConnMaxLifetime(time.Hour) // Recycle connections

	return db, nil
}
