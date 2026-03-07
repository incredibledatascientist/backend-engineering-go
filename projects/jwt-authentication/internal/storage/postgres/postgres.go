package postgres

import (
	"bookstore/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.Config) (*gorm.DB, error) {
	dbCfg := cfg.Postgres

	if dbCfg.Host == "" || dbCfg.Name == "" {
		return nil, fmt.Errorf("[postgres] database configuration incomplete")
	}

	sslMode := dbCfg.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
		sslMode,
		dbCfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Connection Pooling
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	return nil, err
	// }

	// sqlDB.SetMaxOpenConns(25)
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
