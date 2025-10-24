package config

import "os"

type Config struct {
	DatabaseUrl string
	RedisAddr   string
	JWTSecret   string
	Port        string
}

func Load() *Config {
	return &Config{
		DatabaseUrl: GetEnv("DSN", "Nothing"),
		RedisAddr:   GetEnv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:   GetEnv("JWT_SECRET", "1234"),
		Port:        GetEnv("PORT", "8080"),
	}
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
