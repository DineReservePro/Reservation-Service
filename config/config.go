package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_DATABASE string
	DB_PASSWORD string
	GRPC_PORT   string
}

var Logger *slog.Logger

func InitLogger() {
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, nil)
	Logger = slog.New(handler)
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	cfg := Config{}
	cfg.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(Coalesce("DB_PORT", 5432))
	cfg.DB_USERNAME = cast.ToString(Coalesce("DB_USERNAME", "postgres"))
	cfg.DB_DATABASE = cast.ToString(Coalesce("DB_DATABASE", "reservation_service"))
	cfg.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "03212164"))
	cfg.GRPC_PORT = cast.ToString(Coalesce("GRPC_PORT", ":50051"))

	return cfg
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
