package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	PostgresURL       string
	RedisAddr         string
	RedisPassword     string
	CORSAllowedOrigin string
}

func Load() Config {
	_ = godotenv.Load(".env", "backend/.env")

	return Config{
		Port:              getEnv("PORT", "8080"),
		PostgresURL:       getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:5432/app_db?sslmode=disable"),
		RedisAddr:         getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:     getEnv("REDIS_PASSWORD", ""),
		CORSAllowedOrigin: getEnv("CORS_ALLOWED_ORIGIN", "http://localhost:3000"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
