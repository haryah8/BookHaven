package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var DB_HOST string
var DB_PORT string
var DB_NAME string
var DB_USER string
var DB_PASS string

var PORT string
var XENDIT_SECRET_KEY string
var SECRET_KEY string
var NINJA_API_KEY string

func InitConfig(logger *logrus.Logger) {

	err := godotenv.Load("config/.env")
	if err != nil {
		logger.Fatal("FAILED LOAD CONFIG", err)
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	PORT = os.Getenv("PORT")
	XENDIT_SECRET_KEY = os.Getenv("XENDIT_SECRET_KEY")
	NINJA_API_KEY = os.Getenv("NINJA_API_KEY")
	SECRET_KEY = os.Getenv("SECRET_KEY")

	logger.Info("SUCCESS LOAD CONFIG")
}
