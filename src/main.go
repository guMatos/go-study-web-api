package main

import (
	"fmt"
	"net/http"
	"study-webapi/infra"

	"study-webapi/domain"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		println(err.Error())
		return
	}
	port := viper.Get("PORT")
	if port == nil {
		println("Port number has no value")
	}

	router := gin.Default()
	router.GET("/album", getAlbums)
	router.POST("/album", addAlbum)

	environment := fmt.Sprintf("localhost:%v", port)

	router.Run(environment)
}

func getAlbums(c *gin.Context) {
	repo := infra.Repository{}
	albuns, err := repo.GetAlbums()
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
	repo := infra.Repository{}
	var newAlbum domain.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := repo.AddAlbum(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Created"})
}
