package database

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/wallacemachado/api-bank-transfers/src/config"
	"github.com/wallacemachado/api-bank-transfers/src/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func StartDatabase(dbName string) *gorm.DB {
	config.Init()
	fmt.Println("Could not connect to the Postgres Database")
	DbHost := config.DBHost
	DbPort := config.DBPort
	DbUser := config.DBUser
	DbName := dbName
	DbSSlMode := config.DBSslMode
	DbPass := config.DBPass

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", DbHost, DbPort, DbUser, DbName, DbSSlMode, DbPass)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Could not connect to the Postgres Database1")
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database2")
		log.Fatal("Error: ", err)
	}

	db = database

	database.DB()

	migrations.RunAutoMigrations(db)

	if dbName != "test" {
		db.Exec("CREATE DATABASE test")
	}

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
