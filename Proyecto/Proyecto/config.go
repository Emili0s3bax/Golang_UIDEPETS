package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

func LoadConfig() (*Config, error) {

	godotenv.Load()

	config := &Config{
		Server:   getEnv("DB_SERVER", "DESKTOP-SQR4APT\\BDDUIDEPETS"),
		Port:     getEnv("DB_PORT", "1433"),
		Database: getEnv("DB_NAME", "Info"),
		User:     getEnv("DB_USER", "sa"),
		Password: getEnv("DB_PASSWORD", "root"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
