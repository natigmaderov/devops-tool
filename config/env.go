package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Host:     getEnv("HOST", "localhost"),
		Port:     getEnv("PORT", "5432"),
		User:     getEnv("USER", "admin"),
		Password: getEnv("PASSWORD", "admin"),
		DBName:   getEnv("DBNAME", "natig"),
		SSLMode:  getEnv("SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
