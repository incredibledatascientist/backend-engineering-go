package config

import (
	"fmt"
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

// Storage type definition
type StorageType string

const (
	StoragePostgres StorageType = "postgres"
	StorageSQLite   StorageType = "sqlite"
	StorageMemory   StorageType = "memory"
)

// Validate storage type
func (s StorageType) IsValid() bool {
	switch s {
	case StoragePostgres, StorageSQLite, StorageMemory:
		return true
	default:
		return false
	}
}

// TLS configuration
type TLSConfig struct {
	UseTLS   bool   `yaml:"use_tls"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

// HTTP server configuration
type HTTPServer struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
	TLS          TLSConfig     `yaml:"tls"`
}

// JWT configuration
type JWTConfig struct {
	Secret string `yaml:"secret"`
}

// SQLite configuration
type SQLiteConfig struct {
	Name string `yaml:"name"`
}

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
	Env      string         `yaml:"env"`
	Server   HTTPServer     `yaml:"server"`
	Storage  StorageType    `yaml:"storage"`
	JWT      JWTConfig      `yaml:"jwt"`
	Postgres PostgresConfig `yaml:"postgres"`
	SQLite   SQLiteConfig   `yaml:"sqlite"`
}

// Load configuration from YAML file
func LoadConfig(configFile string) (Config, error) {

	// Read config file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	// Parse YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	// Validate storage
	if !cfg.Storage.IsValid() {
		return Config{}, fmt.Errorf("invalid storage type: %s", cfg.Storage)
	}

	return cfg, nil
}
