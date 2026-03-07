package main

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Postgres configuration
type PostgresConfig struct {
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     string `env:"DB_PORT"`
	DatabaseName     string `env:"DB_NAME"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`
	DatabaseTimeZone string `env:"DB_TIMEZONE"`
	DatabaseSSLMode  string `env:"DB_SSL_MODE"`
}

func LoadConfig() (*PostgresConfig, error) {
	cfg := &PostgresConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("[load_config] fialed to load (%v)", err)
	}

	return cfg, nil
}
