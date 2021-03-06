package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	ServerHost   string
	ServerPort   string
	DBHost       string
	DBPort       string
	DBSslMode    string
	DBUser       string
	DBName       string
	DBPass       string
	DBType       string
	SecretKeyJwt string
)

func Init() {

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
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
	SecretKeyJwt = os.Getenv("SECRET_KEY_JWT")

}
