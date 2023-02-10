package service

import (
	"gomock/database"
	"gomock/logger"
	"strconv"
)

const (
	ERROR_ID   = -1
	NIL_TITLE  = ""
	NIL_ARTIST = ""
	NIL_PRICE  = ""
)

// Create new album info and add it to the database
func addAlbumToDB(newAlbum AlbumInfo) (int64, error) {
	logger.Log.SetPrefix("[addAlbumToDB] ")

	dbAlbum := database.NewAlbulm(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	id, err := database.AddAlbum(dbAlbum)

	if err != nil {
		logger.Log.Println(err)
		return ERROR_ID, err
	}

	logger.Log.Println("Album added to database with id: ", id)
	return id, nil
}

// Validate the album id and delete the album from the database
func deleteAlbumFromDB(id int64) error {
	logger.Log.SetPrefix("[deleteAlbumFromDB] ")

	_, err := database.GetAlbumByID(id)

	if err != nil {
		logger.Log.Println(err)
		return err
	}

	err = database.DeleteAlbum(id)

	if err != nil {
		logger.Log.Println(err)
		return err
	}

	logger.Log.Println("Album deleted from database with id: ", id)
	return nil
}

// Validate the album id and return the album from the database
func getAlbumFromDB(id int64) (database.Album, error) {
	logger.Log.SetPrefix("[getAlbumFromDB] ")

	dbAlbum, err := database.GetAlbumByID(id)

	if err != nil {
		logger.Log.Println(err)
		return database.Album{}, err
	}

	logger.Log.Println("Album retrieved from database with id: ", id)
	return dbAlbum, nil
}

// Validate the new album information and apply the changes to the database
func editAlbumFromDB(newAlbumInfo AlbumStr) (database.Album, error) {
	logger.Log.SetPrefix("[editAlbumFromDB] ")

	oldAlbum, err := database.GetAlbumByID(newAlbumInfo.ID)

	if err != nil {
		logger.Log.Println(err)
		return database.Album{}, err
	}

	if newAlbumInfo.Title != NIL_TITLE {
		oldAlbum.Title = newAlbumInfo.Title
	}

	if newAlbumInfo.Artist != NIL_ARTIST {
		oldAlbum.Artist = newAlbumInfo.Artist
	}

	if newAlbumInfo.Price != NIL_PRICE {
		converted_price, err := strconv.ParseFloat(newAlbumInfo.Price, 32)
		if err != nil {
			logger.Log.Println(err)
			return database.Album{}, err
		}

		oldAlbum.Price = float32(converted_price)
	}

	err = database.EditAlbum(oldAlbum.ID, oldAlbum)

	if err != nil {
		logger.Log.Println(err)
		return database.Album{}, err
	}

	logger.Log.Println("Album edited in database with id: ", oldAlbum.ID)
	return oldAlbum, nil
}
