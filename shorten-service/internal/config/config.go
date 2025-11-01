package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl  string
	Port   string
	AppEnv string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}
	cfg := &Config{
		DBUrl:  getenv("DB_URL", ""),
		Port:   getenv("PORT", "8001"),
		AppEnv: getenv("APP_ENV", "development"),
	}
	if cfg.DBUrl == "" {
		log.Fatal("DB_URL not set in environment variables")
	}
	return cfg
}
func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
