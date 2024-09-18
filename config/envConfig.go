package config

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvConfig(key string) string {
	err := godotenv.Load(".env")
	// No need to load env inside container...
	if err != nil {
	}

	return os.Getenv(key)
}
