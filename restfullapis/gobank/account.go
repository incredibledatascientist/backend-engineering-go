package main

import "fmt"

var CUSTOMERS []Account

func GenAccountNumber() string {
	count := len(CUSTOMERS) + 1
	return fmt.Sprintf("%010d", count)
}

func NewAccount(firstName, lastName string, balance float64) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    "55555",
		Balance:   balance,
	}
}
