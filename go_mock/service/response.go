package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// receive the album information from the client and respond with the album id
func PostAlbum(c *gin.Context) {
	newAlbum, receive_err := receivePostAlbum(c)

	if receive_err != nil {
		return
	}

	id, err := AddAlbumToDB(newAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

// receive the album id from the client and respond with the deleted album id
func DeleteAlbum(c *gin.Context) {
	id, receive_err := receiveDeleteAlbumId(c)

	if receive_err != nil {
		return
	}

	err := DeleteAlbumFromDB(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

// receive the album id from the client and respond with the album information
func ReturnAlbum(c *gin.Context) {
	id, receive_err := receiveGetAlbumById(c)

	if receive_err != nil {
		return
	}

	album, err := GetAlbumFromDB(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// receive the album information from the client and respond with the edited album information
func EditAlbum(c *gin.Context) {
	newAlbum, receive_err := receiveEditAlbum(c)

	if receive_err != nil {
		return
	}

	editedAlbum, err := EditAlbumFromDB(newAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, editedAlbum)
}
