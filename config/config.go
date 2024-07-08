package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	GRPC_PORT   string
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_DATABASE string
	DB_PASSWORD string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}

	cfg.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	return cfg
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultValue
}
