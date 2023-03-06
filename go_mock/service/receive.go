package service

import (
	"gomock/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumStr struct {
	ID     int64
	Title  string
	Artist string
	Price  string
}

// Receive the album info from the client and return data with type AlbumInfo
func receivePostAlbum(c *gin.Context) (AlbumInfo, error) {
	var newAlbum AlbumInfo

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.WarnLogger.Printf("Error binding JSON: %s\n", err)
		return newAlbum, err
	}

	logger.InfoLogger.Printf("Receive post album request, album info: %v", newAlbum)

	return newAlbum, nil
}

// Receive the album id from the client and return data with type int64
func receiveDeleteAlbumId(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.WarnLogger.Printf("Error parsing id: %s\n", err)
		return id, err
	}

	logger.InfoLogger.Printf("Receive delete album request, album id: %d", id)

	return id, nil
}

// Receive the album id from the client and return data with type int64
func receiveGetAlbumById(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.WarnLogger.Printf("Error parsing id: %s\n", err)
		return id, err
	}

	logger.InfoLogger.Printf("Receive get album request, album id: %d", id)

	return id, nil
}

// Receive the album info from the client and return data with type AlbumStr
func receiveEditAlbum(c *gin.Context) (AlbumStr, error) {
	var album AlbumStr

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.WarnLogger.Printf("Error binding JSON: %s\n", err)
		return album, err
	}

	logger.InfoLogger.Printf("Receive edit album request, album info: %v", album)

	return album, nil
}
