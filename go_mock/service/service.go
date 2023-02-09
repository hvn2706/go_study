package service

import (
	"github.com/gin-gonic/gin"
)

type AlbumInfo struct {
	Title  string
	Artist string
	Price  float32
}

func RunServer() {
	router := gin.Default()
	router.POST("/albums", PostAlbum)
	router.GET("/albums/:id", ReturnAlbum)
	router.PUT("/albums/edit/", EditAlbum)
	router.DELETE("/albums/delete/:id", DeleteAlbum)

	router.Run(":8080")
}
