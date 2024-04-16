package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnvironments() {
	if os.Getenv("ENVIRONMENT") != "Production" {
		readEnvVars()
	}
}

func readEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
