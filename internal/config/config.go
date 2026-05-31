package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	AppEnv  string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}

	return Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
