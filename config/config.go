package config

import (
	"os"
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
	Config.MYSQL_USERNAME = os.Getenv("DB_USER")
	Config.MYSQL_PASSWORD = os.Getenv("DB_PASSWORD")
	Config.MYSQL_DATABASE = os.Getenv("DB_NAME")
	Config.MYSQL_HOST = os.Getenv("DB_HOST")
	Config.MYSQL_PORT = os.Getenv("DB_PORT")
	Config.MEDIA_HOST_PATH = os.Getenv("MEDIA_HOST_PATH")
}
