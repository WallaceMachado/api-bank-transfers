package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerHost string
	ServerPort string
	DBHost     string
	DBPort     string
	DBSslMode  string
	DBUser     string
	DBName     string
	DBPass     string
	DBType     string
)

func Init() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ServerHost = os.Getenv("HOST")
	ServerPort = os.Getenv("PORT")

	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBSslMode = os.Getenv("DB_SSL_MODE")
	DBUser = os.Getenv("DB_USER")
	DBName = os.Getenv("DB_NAME")
	DBPass = os.Getenv("DB_PASS")
	DBType = os.Getenv("DB_TYPE")

}
