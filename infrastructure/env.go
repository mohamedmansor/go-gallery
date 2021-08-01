package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

// load ENV
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}
}
