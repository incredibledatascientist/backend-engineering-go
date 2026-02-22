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
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    string    `json:"number"`
	Password  string    `json:"_"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// Account Schema
type AccountSchema struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Password  string  `json:"password"`
	Balance   float64 `json:"balance"`
}

// Balance transfer schema
type TransferSchema struct {
	ToAccount string  `json:"to_account"`
	Amount    float64 `json:"amount"`
}

// User
// type User struct {
// 	Username  string
// 	Password  string
// 	FirstName string
// 	LastName  string
// 	email     string
// }

// Login Request
type LoginRequest struct {
	Number   string `json:"username"`
	Password string `json:"password"`
}
