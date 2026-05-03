package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadOnce sync.Once

func LoadEnv() {
	loadOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println(".env file not found, using system environment")
		}
	})
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}