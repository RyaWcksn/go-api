package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }
    value := os.Getenv(key)
    return value
}
