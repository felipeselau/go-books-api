package server

import (
	"log"

	"github.com/felipeselau/go-books-api/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func Run(s *Server) {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server running on port " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
