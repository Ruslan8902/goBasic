package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key     string
	BaseUrl string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Key:     os.Getenv("KEY"),
		BaseUrl: "https://api.jsonbin.io/v3/b",
	}
}
