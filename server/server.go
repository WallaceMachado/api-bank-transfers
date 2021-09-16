package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/config"
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
	fmt.Println(s.port)
	log.Printf("Server running at port: %s", s.port)

	gin.Default().Run(":" + s.port)

}
