package utils

import (
	"time"

	"github.com/ahmadammarm/courses-management/backend/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET"))

func JWTGenerate(username string) string {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return ""
	}

	return token
}
