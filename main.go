package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/album", getAlbums)
	router.POST("/album", addAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	repo := repository{}
	albuns, err := repo.getAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	if albuns == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.IndentedJSON(http.StatusOK, albuns)
}

func addAlbum(c *gin.Context) {
	repo := repository{}
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := repo.addAlbum(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Created"})
}
