package main

import (
	"github.com/wallacemachado/api-bank-transfers/config"
	"github.com/wallacemachado/api-bank-transfers/database"
	"github.com/wallacemachado/api-bank-transfers/server"
)

func main() {
	config.Init()

	database.StartDatabase()

	server := server.NewServer()
	server.Run()

}
