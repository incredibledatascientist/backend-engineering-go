package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig holds all configurations
type AppConfig struct {
	ServerPort string
	DBType     string // "postgres" or "memory"
	JWTSecret  string
	HashCost   int
	Postgres   PostgresConfig
}

// PostgresConfig holds the DB credentials
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

// LoadConfig reads variables from .env or OS environment
func LoadConfig() *AppConfig {
	err := GodotenvLoad()
	if err != nil {
		log.Println("Note: No .env file found, relying on system environment variables")
	}

	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	hashCost, _ := strconv.Atoi(getEnv("HASH_COST", "10"))

	return &AppConfig{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBType:     getEnv("DB_TYPE", "postgres"),
		JWTSecret:  getEnv("JWT_SECRET", "master@golang"),
		HashCost:   hashCost,
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "infierms"),
			Name:     getEnv("DB_NAME", "ginjwtauth"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "Asia/Kolkata"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GodotenvLoad wrap so we don't import joho wrongly in raw form usually
func GodotenvLoad() error {
	return godotenv.Load()
}
