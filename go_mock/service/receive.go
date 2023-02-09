package service

import (
	"gomock/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func receivePostAlbum(c *gin.Context) (AlbumInfo, error) {
	logger.Log.SetPrefix("[receivePostAlbum] ")
	var newAlbum AlbumInfo

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Log.Println(err)
		return newAlbum, err
	}

	return newAlbum, nil
}

func receiveDeleteAlbumId(c *gin.Context) (int64, error) {
	logger.Log.SetPrefix("[receiveGetAlbumById] ")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Log.Println(err)
		return id, err
	}

	return id, nil
}

func receiveGetAlbumById(c *gin.Context) (int64, error) {
	logger.Log.SetPrefix("[receiveGetAlbumById] ")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Log.Println(err)
		return id, err
	}

	return id, nil
}
