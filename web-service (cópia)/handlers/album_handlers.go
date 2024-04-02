// handlers/album_handlers.go
package handlers

import (
	"net/http"

	"./models"

	"github.com/gin-gonic/gin"
)

// GetAlbums retorna a lista de todos os álbuns como JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

// PostAlbums adiciona um álbum recebido em JSON no corpo da solicitação.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// PutAlbumsByID atualiza um álbum existente com o ID fornecido.
func PutAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum models.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DeleteAlbumsByID exclui um álbum existente com o ID fornecido.
func DeleteAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums = append(models.Albums[:i], models.Albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, models.Albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
