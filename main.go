package main

import (
	"github.com/wallacemachado/api-bank-transfers/src/config"
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/server"
)

func main() {
	config.Init()

	database.StartDatabase()

	server := server.NewServer()
	server.Run()

}
