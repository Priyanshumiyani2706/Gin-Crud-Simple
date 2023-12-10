package Config

import (
	"log"

	"github.com/lpernett/godotenv"
)

func LoadEnv() {
	err := godotenv.Load() // load environment variable file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
