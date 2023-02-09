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

func addAlbumToDB(newAlbum AlbumInfo) (int64, error) {
	logger.Log.SetPrefix("[addAlbumToDB] ")

	dbAlbum := database.NewAlbulm(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	id, err := database.AddAlbum(dbAlbum)

	if err != nil {
		logger.Log.Println(err)
		return ERROR_ID, err
	}

	return id, nil
}

func deleteAlbumFromDB(id int64) error {
	logger.Log.SetPrefix("[deleteAlbumFromDB] ")

	err := database.DeleteAlbum(id)

	if err != nil {
		logger.Log.Println(err)
		return err
	}

	return nil
}

func getAlbumFromDB(id int64) (database.Album, error) {
	logger.Log.SetPrefix("[getAlbumFromDB] ")

	dbAlbum, err := database.GetAlbumByID(id)

	if err != nil {
		logger.Log.Println(err)
		return database.Album{}, err
	}

	return dbAlbum, nil
}

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

	return oldAlbum, nil
}
