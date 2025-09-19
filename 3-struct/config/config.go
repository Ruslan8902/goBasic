package config

import (
	"fmt"
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
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Key:     os.Getenv("KEY"),
		BaseUrl: "https://api.jsonbin.io/v3/b",
	}
}
