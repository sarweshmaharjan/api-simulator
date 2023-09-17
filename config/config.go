package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	APIURL string
	Port   string
	// Add more configuration fields here
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &Config{
		APIURL: os.Getenv("API_URL"),
		Port:   os.Getenv("PORT"),
		// Initialize other configuration fields here
	}

	return config, nil
}
