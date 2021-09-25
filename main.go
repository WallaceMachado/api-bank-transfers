package main

import (
	"github.com/wallacemachado/api-bank-transfers/src/server"
	"github.com/wallacemachado/api-bank-transfers/src/shared/config"
	"github.com/wallacemachado/api-bank-transfers/src/shared/database"
)

func main() {
	config.Init()

	database.StartDatabase(config.DBName)

	server := server.NewServer()
	server.Run()

}
