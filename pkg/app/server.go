package app

import (
	"go-web/pkg/api"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	albumService api.AlbumService
}

func NewServer(router *gin.Engine, albumService api.AlbumService) *Server {
	return &Server{
		router:       router,
		albumService: albumService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run(":8888")

	if err != nil {
		log.Printf("Failed to run router: %v", err)
		return err
	}

	return nil
}
