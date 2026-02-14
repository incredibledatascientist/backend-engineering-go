package main

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
	Store  Storage
}

// Account Details
type Account struct {
	Id        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Number    string  `json:"number"`
	Balance   float64 `json:"balance"`
}
