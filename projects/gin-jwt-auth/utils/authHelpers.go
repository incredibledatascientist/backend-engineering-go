package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// func CheckUserRole(c *gin.Context, role models.UserRole) error {

// }

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(hashPassword), err
}

func VerifyPassword(userPassword, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return false, err.Error()
	}

	return true, "User verification successful."
}
