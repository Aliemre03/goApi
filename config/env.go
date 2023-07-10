package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error env load %v ", err)
	}
	return os.Getenv("MONGOURI")
}
