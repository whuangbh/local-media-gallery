package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config struct {
	MYSQL_USERNAME  string
	MYSQL_PASSWORD  string
	MYSQL_DATABASE  string
	MYSQL_HOST      string
	MYSQL_PORT      string
	MEDIA_HOST_PATH string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Config.MYSQL_USERNAME = os.Getenv("MYSQL_USERNAME")
	Config.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	Config.MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
	Config.MYSQL_HOST = os.Getenv("MYSQL_HOST")
	Config.MYSQL_PORT = os.Getenv("MYSQL_PORT")
	Config.MEDIA_HOST_PATH = os.Getenv("MEDIA_HOST_PATH")
}
