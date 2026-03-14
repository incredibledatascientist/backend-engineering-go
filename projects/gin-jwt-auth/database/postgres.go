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

// Neon DB Postgres configuration
type NeonDB struct {
	ConnString string
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

func InitDatabase() (*gorm.DB, error) {
	// postgres := PostgresConfig{
	// 	Name:     "ginjwtauth",
	// 	Port:     5432,
	// 	Host:     "localhost",
	// 	User:     "postgres",
	// 	Password: "infierms",
	// 	TimeZone: "Asia/Kolkata",
	// 	SSLMode:  "disable",
	// }

	// Neon DB configs
	// neonDB := "postgresql://user:password@host/dbname?sslmode=require&channel_binding=require"
	postgres := PostgresConfig{
		Name:     "neondb",
		Port:     5432,
		Host:     "ep-wispy-heart-ams3cs64-pooler.c-5.us-east-1.aws.neon.tech",
		User:     "neondb_owner",
		Password: "npg_xAHyU0X6aMLI",
		TimeZone: "Asia/Kolkata",
		SSLMode:  "require",
	}

	db, err := NewPostgresDB(Config{
		Postgres: postgres,
	})

	if err != nil {
		return nil, err
	}

	initDB = db // for temp

	return db, nil
}

var initDB *gorm.DB

func GetDB() *gorm.DB {
	return initDB
}
