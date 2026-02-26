package models

import (
	"time"
)

// Server Configurations
type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// TLS/SSL Configs
type TLSConfig struct {
	UseTLS   bool
	CertFile string
	KeyFile  string
}

// APIServer Details
type APIServer struct {
	Name   string       `json:"name"`
	Addr   string       `json:"addr"`
	Server ServerConfig `json:"server"`
	TLS    TLSConfig    `json:"tls"`
	// Store  Storage
}
