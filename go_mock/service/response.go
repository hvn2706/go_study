package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAlbum(c *gin.Context) {
	newAlbum, receive_err := receivePostAlbum(c)

	if receive_err != nil {
		return
	}

	id, err := addAlbumToDB(newAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func DeleteAlbum(c *gin.Context) {
	id, receive_err := receiveDeleteAlbumId(c)

	if receive_err != nil {
		return
	}

	err := deleteAlbumFromDB(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func ReturnAlbum(c *gin.Context) {
	id, receive_err := receiveGetAlbumById(c)

	if receive_err != nil {
		return
	}

	album, err := getAlbumFromDB(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func EditAlbum(c *gin.Context) {
	newAlbum, receive_err := receiveEditAlbum(c)

	if receive_err != nil {
		return
	}

	editedAlbum, err := editAlbumFromDB(newAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, editedAlbum)
}
