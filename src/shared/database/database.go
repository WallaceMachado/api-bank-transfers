package database

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/wallacemachado/api-bank-transfers/src/shared/config"
	"github.com/wallacemachado/api-bank-transfers/src/shared/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func StartDatabase(dbName string) *gorm.DB {
	config.Init()

	DbHost := config.DBHost
	DbPort := config.DBPort
	DbUser := config.DBUser
	DbName := dbName
	DbSSlMode := config.DBSslMode
	DbPass := config.DBPass

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", DbHost, DbPort, DbUser, DbName, DbSSlMode, DbPass)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("Error: ", err)
	}

	db = database

	database.DB()

	migrations.RunAutoMigrations(db)

	fmt.Println("Connect to Database!")

	return db

}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func SetDatabase(newDb *gorm.DB) {
	db = newDb
}
