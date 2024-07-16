package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment")
	}
}
