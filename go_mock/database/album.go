package database

import (
	"fmt"
	"gomock/logger"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func NewAlbulm(title, artist string, price float64) Album {
	return Album{
		Title:  title,
		Artist: artist,
		Price:  price,
	}
}

func GetAlbumByID(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		logger.WarnLogger.Printf("Error getting album with id=%d from database: %s\n", id, err)
		return Album{}, fmt.Errorf("getAlbumByID %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		logger.WarnLogger.Printf("Error adding album to database: %s\n", err)
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.ErrorLogger.Printf("Error getting last inserted album id to database: %s\n", err)
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func DeleteAlbum(id int64) error {
	result, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		logger.WarnLogger.Printf("Error deleting album with id=%d from database: %s\n", id, err)
		return fmt.Errorf("deleteAlbum: %v", err)
	}

	rows_affected, err := result.RowsAffected()

	if err != nil {
		logger.ErrorLogger.Printf("Error getting rows affected by delete album with id=%d from database: %s\n", id, err)
		return fmt.Errorf("deleteAlbum: %v", err)
	}

	if rows_affected == 0 {
		return fmt.Errorf("deleteAlbum: no rows affected, album with id %d might not exist", id)
	}
	return nil
}

func EditAlbum(alb Album) error {
	result, err := db.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		logger.WarnLogger.Printf("Error editing album with id=%d from database: %s\n", alb.ID, err)
		return fmt.Errorf("editAlbum: %v", err)
	}

	rows_affected, err := result.RowsAffected()

	if err != nil {
		logger.ErrorLogger.Printf("Error getting rows affected by edit album with id=%d from database: %s\n", alb.ID, err)
		return fmt.Errorf("editAlbum: %v", err)
	}

	if rows_affected == 0 {
		logger.InfoLogger.Printf("No changes apply by edit album action with id=%d from database\n", alb.ID)
	}

	return nil
}
