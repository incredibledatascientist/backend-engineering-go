package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var CUSTOMERS []Account

func GenAccountNumber() string {
	count := len(CUSTOMERS) + 1
	return fmt.Sprintf("%010d", count)
}

func (acc Account) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
}

func NewAccount(firstName, lastName string, password string, balance float64) (*Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	account := &Account{
		FirstName: firstName,
		LastName:  lastName,
		Password:  string(hashedPassword),
		Balance:   balance,
	}
	return account, nil
}
