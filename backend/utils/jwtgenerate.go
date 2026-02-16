package utils

import (
	"time"

	"github.com/ahmadammarm/courses-management/backend/config"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	InstructorID string `json:"instructor_id"`
	Username     string `json:"username"`
	jwt.RegisteredClaims
}

func JWTGenerate(instructorID string, username string) (string, error) {

	jwtKey := []byte(config.GetEnv("JWT_SECRET"))

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &CustomClaims{
		InstructorID: instructorID,
		Username:     username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "courses-management-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
