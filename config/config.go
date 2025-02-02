package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string

	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_REDIRECT_URL  string

	JWT_SECRET string
)

func Init() {
	fmt.Println("Initializing config")
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found")
	}

	DATABASE_USERNAME = getEnvOrDefault("DATABASE_USERNAME", "postgres")
	DATABASE_PASSWORD = getEnvOrDefault("DATABASE_PASSWORD", "postgres")
	DATABASE_HOST = getEnvOrDefault("DATABASE_HOST", "localhost")
	DATABASE_PORT = getEnvOrDefault("DATABASE_PORT", "5432")
	DATABASE_NAME = getEnvOrDefault("DATABASE_NAME", "lily_db")

	GOOGLE_CLIENT_ID = getEnvOrDefault("GOOGLE_CLIENT_ID", "")
	GOOGLE_CLIENT_SECRET = getEnvOrDefault("GOOGLE_CLIENT_SECRET", "")
	GOOGLE_REDIRECT_URL = getEnvOrDefault("GOOGLE_REDIRECT_URL", "")

	JWT_SECRET = getEnvOrDefault("JWT_SECRET", "")
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
