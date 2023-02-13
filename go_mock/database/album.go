package database

import (
	"fmt"
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

func FindAlbumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func GetAlbumByID(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		return Album{}, fmt.Errorf("getAlbumByID %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func DeleteAlbum(id int64) error {
	_, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleteAlbum: %v", err)
	}
	return nil
}

func EditAlbum(id int64, alb Album) error {
	_, err := db.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price, id)
	if err != nil {
		return fmt.Errorf("editAlbum: %v", err)
	}
	return nil
}
