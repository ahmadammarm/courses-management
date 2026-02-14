package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// load environment variables from .env file, if it exists.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}
}

// get environment variable by key, return empty string if not found
func GetEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return ""
}
