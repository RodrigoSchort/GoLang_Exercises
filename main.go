package main

import (
	"./handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.PUT("/albums/:id", handlers.PutAlbumsByID)
	router.DELETE("/albums/:id", handlers.DeleteAlbumsByID)

	router.Run("localhost:8080")
}
