package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher defines behavior for hashing passwords
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashed, password string) bool
}

type bcryptHasher struct {
	cost int
}

// NewBcryptHasher creates a Bcrypt implementation of PasswordHasher
func NewBcryptHasher(cost ...int) PasswordHasher {
	hashCost := bcrypt.DefaultCost
	if len(cost) > 0 {
		hashCost = cost[0]
	}
	return &bcryptHasher{cost: hashCost}
}

func (h *bcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (h *bcryptHasher) Compare(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
