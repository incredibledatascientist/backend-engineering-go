package main

import "fmt"

type Account struct {
	Id        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Number    string  `json:"number"`
	Balance   float64 `json:"balance"`
}

var CUSTOMERS []Account

func GenId() int64 {
	return int64(len(CUSTOMERS) + 1)
}

func GenAccountNumber() string {
	count := len(CUSTOMERS) + 1
	return fmt.Sprintf("%010d", count)
}

func NewAccount(firstName, lastName string, balance float64) *Account {
	return &Account{
		Id:        GenId(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    GenAccountNumber(),
		Balance:   balance,
	}
}
