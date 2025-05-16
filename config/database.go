package config

import (
	"os"
)

type Config struct {
	DatabaseUrl string
	Port        string
}

func Load() *Config {
	return &Config{
		DatabaseUrl: getEnv("DATABASE_URL", ""),
		Port:        getEnv("PORT", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
