package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	Username string
	Role     string
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "master@golang" // fallback
	}
	return []byte(secret)
}

func GenerateAuthTokens(userId, username, role string) (string, string) {
	claims := &SignedDetails{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Short lived
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId,
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)), // 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(getJWTSecret())
	if err != nil {
		log.Fatal(err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(getJWTSecret())
	if err != nil {
		log.Fatal(err)
	}

	return accessToken, refreshToken
}
