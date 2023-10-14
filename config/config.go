package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Cfg config

type config struct {
	APIURL           string
	Port             string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	// Add more configuration fields here
}

func LoadConfig() error {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Cfg = config{
		APIURL:           os.Getenv("API_URL"),
		Port:             os.Getenv("PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		// Initialize other configuration fields here
	}
	return nil
}
