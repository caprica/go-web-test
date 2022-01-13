package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "album API ok",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetArtistAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		c.Header("Content-Type", "application/json")

		result, err := s.albumService.ArtistAlbumNames(name)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		if result != nil {
			c.JSON(http.StatusOK, result)
		} else {
			c.String(http.StatusNotFound, "Nothing found for "+name)
		}
	}
}
