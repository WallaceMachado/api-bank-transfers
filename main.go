package main

import (
	"github.com/wallacemachado/api-bank-transfers/server"
)

func main() {
	server := server.NewServer()
	server.Run()

}
