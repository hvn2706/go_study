package service

import (
	"errors"
	"gomock/database"
	"gomock/logger"
	"strconv"
)

const (
	ERROR_ID             = -1
	NIL_TITLE            = ""
	NIL_ARTIST           = ""
	MIN_PRICE_NUM        = 0
	NIL_PRICE_STR        = ""
	ACCEPTED_PRICE_ERROR = 0.01
)

// Create new album info and add it to the database
func AddAlbumToDB(newAlbum AlbumInfo) (int64, error) {
	if newAlbum.Title == NIL_TITLE || newAlbum.Artist == NIL_ARTIST || newAlbum.Price < MIN_PRICE_NUM {
		logger.InfoLogger.Println("Invalid album information, NIL value or negative price")
		return ERROR_ID, errors.New("invalid album information, NIL value or negative price")
	}

	dbAlbum := database.NewAlbulm(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	id, err := database.AddAlbum(dbAlbum)

	if err != nil {
		logger.ErrorLogger.Printf("Error adding album to database: %s\n", err)
		return ERROR_ID, err
	}

	logger.InfoLogger.Printf("Album added to database with id: %d\n", id)
	return id, nil
}

// Validate the album id and delete the album from the database
func DeleteAlbumFromDB(id int64) error {
	err := database.DeleteAlbum(id)

	if err != nil {
		logger.InfoLogger.Printf("Error deleting album with id=%d from database: %s\n", id, err)
		return err
	}

	logger.InfoLogger.Printf("Album deleted from database with id: %d\n", id)
	return nil
}

// Validate the album id and return the album from the database
func GetAlbumFromDB(id int64) (database.Album, error) {
	dbAlbum, err := database.GetAlbumByID(id)

	if err != nil {
		logger.InfoLogger.Printf("Album with id: %d does not exist in database\n", id)
		return database.Album{}, err
	}

	logger.InfoLogger.Printf("Album retrieved from database with id: %d\n", id)
	return dbAlbum, nil
}

// Validate the new album information and apply the changes to the database
func EditAlbumFromDB(newAlbumInfo AlbumStr) (database.Album, error) {
	oldAlbum, err := database.GetAlbumByID(newAlbumInfo.ID)

	if err != nil {
		logger.InfoLogger.Printf("Album with id: %d does not exist in database\n", newAlbumInfo.ID)
		return database.Album{}, err
	}

	if newAlbumInfo.Title != NIL_TITLE {
		oldAlbum.Title = newAlbumInfo.Title
	}

	if newAlbumInfo.Artist != NIL_ARTIST {
		oldAlbum.Artist = newAlbumInfo.Artist
	}

	if newAlbumInfo.Price != NIL_PRICE_STR {
		converted_price, err := strconv.ParseFloat(newAlbumInfo.Price, 64)
		if err != nil {
			logger.TraceLogger.Printf("Error converting price: %s\n", err)
			return database.Album{}, err
		}

		if converted_price < MIN_PRICE_NUM {
			logger.InfoLogger.Println("Invalid album information, negative price")
			return database.Album{}, errors.New("invalid album information, negative price")
		}

		oldAlbum.Price = float64(converted_price)
	}

	err = database.EditAlbum(oldAlbum)

	if err != nil {
		logger.ErrorLogger.Printf("Error editing album in database: %s\n", err)
		return database.Album{}, err
	}

	logger.InfoLogger.Printf("Album edited in database with id: %d; old value: %v; new request value: %v\n", oldAlbum.ID, oldAlbum, newAlbumInfo)

	return oldAlbum, nil
}
