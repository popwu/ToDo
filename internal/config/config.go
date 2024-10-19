package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
	JWTSecret     string
}

func Load() (*Config, error) {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "todo.db"),
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
