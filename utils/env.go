package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal("ENV: " + key + " not found")
	}
	return value
}
