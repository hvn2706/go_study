package service

import (
	"gomock/database"
	"gomock/logger"
)

const (
	ERROR_ID  = -1
	NIL_PRICE = -1
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

func editAlbumFromDB(id int64, newAlbum AlbumInfo) (database.Album, error) {
	logger.Log.SetPrefix("[editAlbumFromDB] ")

	oldAlbum, err := database.GetAlbumByID(id)

	if err != nil {
		logger.Log.Println(err)
		return database.Album{}, err
	}

	if newAlbum.Title != "" {
		oldAlbum.Title = newAlbum.Title
	}

	if newAlbum.Artist != "" {
		oldAlbum.Artist = newAlbum.Artist
	}

	if newAlbum.Price != NIL_PRICE {
		oldAlbum.Price = newAlbum.Price
	}

	err = database.EditAlbum(oldAlbum.ID, oldAlbum)

	if err != nil {
		logger.Log.Println(err)
		return database.Album{}, err
	}

	return oldAlbum, nil
}
