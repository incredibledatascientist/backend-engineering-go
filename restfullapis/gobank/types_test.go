// Command	What It Tests	Subfolders?	Shows Test Names?
// go test .	Current folder only	No	No
// go test ./... -v	All folders recursively Yes Yes

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account, err := NewAccount("Abhishek", "Kumar", "Boss@123", 5000)
	assert.Nil(t, err)

	fmt.Println("Testing: acc-", account)
}
