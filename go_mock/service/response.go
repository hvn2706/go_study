package service

import (
	"gomock/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// receive the album information from the client and respond with the album id
func PostAlbum(c *gin.Context) {
	logger.Log.SetPrefix("[PostAlbum] ")
	newAlbum, receive_err := receivePostAlbum(c)

	if receive_err != nil {
		logger.Log.Printf("Error receive request post album: %s\n", receive_err)
		return
	}

	id, err := AddAlbumToDB(newAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Log.Printf("Error adding album to DB: %s\n", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

// receive the album id from the client and respond with the deleted album id
func DeleteAlbum(c *gin.Context) {
	logger.Log.SetPrefix("[DeleteAlbum] ")
	id, receive_err := receiveDeleteAlbumId(c)

	if receive_err != nil {
		logger.Log.Printf("Error receive request delete album: %s\n", receive_err)
		return
	}

	err := DeleteAlbumFromDB(id)

	if err != nil {
		logger.Log.Printf("Error deleting album from DB: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

// receive the album id from the client and respond with the album information
func ReturnAlbum(c *gin.Context) {
	logger.Log.SetPrefix("[ReturnAlbum] ")
	id, receive_err := receiveGetAlbumById(c)

	if receive_err != nil {
		logger.Log.Printf("Error receive request return album: %s\n", receive_err)
		return
	}

	album, err := GetAlbumFromDB(id)

	if err != nil {
		logger.Log.Printf("Error getting album from DB: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// receive the album information from the client and respond with the edited album information
func EditAlbum(c *gin.Context) {
	logger.Log.SetPrefix("[EditAlbum] ")
	newAlbum, receive_err := receiveEditAlbum(c)

	if receive_err != nil {
		logger.Log.Printf("Error receive request edit album: %s\n", receive_err)
		return
	}

	editedAlbum, err := EditAlbumFromDB(newAlbum)

	if err != nil {
		logger.Log.Printf("Error editing album from DB: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, editedAlbum)
}
