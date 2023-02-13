package service

import (
	"errors"
	"gomock/database"
	"gomock/logger"
	"math"
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
	logger.Log.SetPrefix("[addAlbumToDB] ")

	if newAlbum.Title == NIL_TITLE {
		logger.Log.Println("Invalid album information, NIL title")
		return ERROR_ID, errors.New("invalid album information, NIL title")
	}

	if newAlbum.Artist == NIL_ARTIST {
		logger.Log.Println("Invalid album information, NIL artist")
		return ERROR_ID, errors.New("invalid album information, NIL artist")
	}

	if newAlbum.Price < MIN_PRICE_NUM {
		logger.Log.Println("Invalid album information, negative price")
		return ERROR_ID, errors.New("invalid album information, negative price")
	}

	dbAlbum := database.NewAlbulm(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	id, err := database.AddAlbum(dbAlbum)

	if err != nil {
		logger.Log.Printf("Error adding album to database: %s\n", err)
		return ERROR_ID, err
	}

	logger.Log.Printf("Album added to database with id: %d\n", id)
	return id, nil
}

// Validate the album id and delete the album from the database
func DeleteAlbumFromDB(id int64) error {
	logger.Log.SetPrefix("[deleteAlbumFromDB] ")

	_, err := database.GetAlbumByID(id)

	if err != nil {
		logger.Log.Printf("Album with id: %d does not exist in database\n", id)
		return err
	}

	err = database.DeleteAlbum(id)

	if err != nil {
		logger.Log.Printf("Error deleting album from database: %s\n", err)
		return err
	}

	logger.Log.Printf("Album deleted from database with id: %d\n", id)
	return nil
}

// Validate the album id and return the album from the database
func GetAlbumFromDB(id int64) (database.Album, error) {
	logger.Log.SetPrefix("[getAlbumFromDB] ")

	dbAlbum, err := database.GetAlbumByID(id)

	if err != nil {
		logger.Log.Printf("Album with id: %d does not exist in database\n", id)
		return database.Album{}, err
	}

	logger.Log.Printf("Album retrieved from database with id: %d\n", id)
	return dbAlbum, nil
}

// Validate the new album information and apply the changes to the database
func EditAlbumFromDB(newAlbumInfo AlbumStr) (database.Album, error) {
	logger.Log.SetPrefix("[editAlbumFromDB] ")

	oldAlbum, err := database.GetAlbumByID(newAlbumInfo.ID)

	if err != nil {
		logger.Log.Printf("Album with id: %d does not exist in database\n", newAlbumInfo.ID)
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
			logger.Log.Printf("Error converting price: %s\n", err)
			return database.Album{}, err
		}

		if converted_price < MIN_PRICE_NUM {
			logger.Log.Println("Invalid album information, negative price")
			return database.Album{}, errors.New("invalid album information, negative price")
		}

		oldAlbum.Price = float64(converted_price)
	}

	err = database.EditAlbum(oldAlbum.ID, oldAlbum)

	if err != nil {
		logger.Log.Printf("Error editing album in database: %s\n", err)
		return database.Album{}, err
	}

	currentAlbum, err := database.GetAlbumByID(oldAlbum.ID)

	if err != nil {
		logger.Log.Printf("Album with id: %d does not exist in database\n", oldAlbum.ID)
		return database.Album{}, err
	}

	if currentAlbum.Title != oldAlbum.Title {
		logger.Log.Printf("Error editing album in database, title does not match: %s, %s\n", currentAlbum.Title, oldAlbum.Title)
		return database.Album{}, errors.New("error editing album in database, title does not match")
	}

	if currentAlbum.Artist != oldAlbum.Artist {
		logger.Log.Printf("Error editing album in database, artist does not match: %s, %s\n", currentAlbum.Artist, oldAlbum.Artist)
		return database.Album{}, errors.New("error editing album in database, artist does not match")
	}

	if math.Abs(currentAlbum.Price-oldAlbum.Price) > 0.01 {
		logger.Log.Printf("Error editing album in database, price does not match: %f, %f\n", currentAlbum.Price, oldAlbum.Price)
		return database.Album{}, errors.New("error editing album in database, price does not match")
	}

	logger.Log.Printf("Album edited in database with id: %d; old value: %s, %s, %f; new request value: %s, %s, %s\n", oldAlbum.ID, oldAlbum.Title, oldAlbum.Artist, oldAlbum.Price, newAlbumInfo.Title, newAlbumInfo.Artist, newAlbumInfo.Price)

	return currentAlbum, nil
}
