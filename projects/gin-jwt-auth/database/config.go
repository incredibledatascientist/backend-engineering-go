package database

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from .env file
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return Config{
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "Asia/Kolkata"),
		},
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
