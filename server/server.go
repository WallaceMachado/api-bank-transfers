package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/config"
	"github.com/wallacemachado/api-bank-transfers/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   config.ServerPort,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.Config(s.server)

	log.Printf("Server running at port: %s", s.port)

	log.Fatal(router.Run(":" + s.port))

}
