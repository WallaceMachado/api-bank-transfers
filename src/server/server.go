package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/server/routes"
	"github.com/wallacemachado/api-bank-transfers/src/shared/config"
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
