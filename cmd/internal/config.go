package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}

	return appConfig
}
