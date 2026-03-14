package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken = errors.New("invalid or expired token")
)

// Claims structure extending standard JWT RegisteredClaims
type Claims struct {
	Username string
	Role     string
	jwt.RegisteredClaims
}

// TokenMaker defines the behavior for creating and verifying JWTs
type TokenMaker interface {
	GenerateTokens(userID, username, role string) (accessToken, refreshToken string, err error)
	VerifyToken(tokenString string) (*Claims, error)
}

type jwtMaker struct {
	secretKey string
}

// NewJWTMaker creates an instance of a JWT maker
func NewJWTMaker(secretKey string) TokenMaker {
	return &jwtMaker{secretKey: secretKey}
}

func (m *jwtMaker) GenerateTokens(userID, username, role string) (string, string, error) {
	accessClaims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // 15m access
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	refreshClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7d refresh
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(m.secretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(m.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (m *jwtMaker) VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
