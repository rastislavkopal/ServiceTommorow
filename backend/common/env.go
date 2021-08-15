package common

import (
	"log"

	"github.com/joho/godotenv"
)

// loads env variables from root .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}
