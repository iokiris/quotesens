package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort   string
	PostgresUser string
	PostgresPass string
	PostgresDB   string
	PostgresHost string
}

func Load() *Config {
	_ = godotenv.Load(".env")

	return &Config{
		ServerPort:   getEnv("PORT", "8080"),
		PostgresUser: getEnv("POSTGRES_USER", "user"),
		PostgresPass: getEnv("POSTGRES_PASSWORD", "pass"),
		PostgresDB:   getEnv("POSTGRES_DB", "quotes"),
		PostgresHost: getEnv("POSTGRES_HOST", "localhost"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
