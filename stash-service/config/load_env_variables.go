package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}

// LoadEnvVariables is responsible for loading env variables from the local .env file.
func LoadEnvVariables(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("env variables loaded")
}

func GetPostgresConfig() *DatabaseConfig{
	return &DatabaseConfig{
		Host: os.Getenv("PG_HOST"),
		Port: 6543,
		User: os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Name: os.Getenv("PG_DB_NAME"),
	}
}