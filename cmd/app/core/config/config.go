package core_config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  int
	DatabaseURL string
}

func LoadConfig() (Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	// Read SERVER_PORT
	portStr := os.Getenv("SERVER_PORT")
	if portStr == "" {
		log.Fatal("SERVER_PORT is not set in environment variables or .env file")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid SERVER_PORT: %s", portStr)
	}

	// Read DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in environment variables or .env file")
	}

	return Config{
		ServerPort:  port,
		DatabaseURL: dbURL,
	}, nil
}
