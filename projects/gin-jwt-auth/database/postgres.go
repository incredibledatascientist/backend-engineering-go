package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres configuration
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	TimeZone string `yaml:"timezone"`
	SSLMode  string `yaml:"ssl_mode"`
}

// Application configuration
type Config struct {
	// Env      string         `yaml:"env"`
	// Server   HTTPServer     `yaml:"server"`
	// Storage  StorageType    `yaml:"storage"`
	// JWT      JWTConfig      `yaml:"jwt"`
	// SQLite   SQLiteConfig   `yaml:"sqlite"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// SingularTable: true -> user [not users]
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
	})
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
