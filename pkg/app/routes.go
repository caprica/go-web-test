package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	api := router.Group("/api")
	{
		api.GET("/status", s.ApiStatus())

		artist := api.Group("/artist/:name")
		{
			artist.GET("", s.GetArtistAlbums())
		}
	}

	return router
}
