package service_test

import (
	"gomock/database"
	"gomock/service"
	"testing"
)

// Test the addAlbumToDB function
func TestAddAlbumToDB(t *testing.T) {
	// Connect to the database
	database.Connectdb()
	defer database.CloseDB()
	// Create a new album
	newAlbums := []service.AlbumInfo{
		{
			Title:  "Test Title",
			Artist: "Test Artist",
			Price:  999999.99,
		},
		{
			Title:  "Test Title 2",
			Artist: "Test Artist 2",
			Price:  20,
		},
		{
			Title:  "",
			Artist: "",
			Price:  -1,
		},
		{
			Title:  "",
			Artist: "",
			Price:  1282.3125,
		},
		{
			Title:  "a",
			Artist: "b",
			Price:  0,
		},
		{
			Title:  "235235",
			Artist: "",
			Price:  -1,
		},
		{
			Title:  "",
			Artist: "qwi8jshf",
			Price:  -1,
		},
		{
			Title:  "˙ˆ˙¨˙ˆ˙",
			Artist: "˙ˆ˙©ƒ†¥˙",
			Price:  0.00000000001,
		},
	}

	for _, newAlbum := range newAlbums {
		// Add the album to the database
		id, _ := service.AddAlbumToDB(newAlbum)

		if newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price < 0 {
			if id != service.ERROR_ID {
				t.Fatalf("Returned ID is not ERROR_ID %d; for album: %v\n", id, newAlbum)
				continue
			} else {
				continue
			}
		}

		// Get the album from the database
		dbAlbum, err := database.GetAlbumByID(id)

		// Check for errors
		if err != nil {
			t.Fatalf("Error getting album from database: %s; Album: %v\n", err, newAlbum)
			continue
		}

		// Check if the album was added correctly
		if dbAlbum.Title != newAlbum.Title {
			t.Fatalf("Album was not added correctly to the database. WRONG TITLE. Expected: %s, Got: %s\n", newAlbum.Title, dbAlbum.Title)
			continue
		}

		if dbAlbum.Artist != newAlbum.Artist {
			t.Fatalf("Album was not added correctly to the database. WRONG ARTIST. Expected: %s, Got: %s\n", newAlbum.Artist, dbAlbum.Artist)
			continue
		}

		if dbAlbum.Price-newAlbum.Price > service.ACCEPTED_PRICE_ERROR {
			t.Fatalf("Album was not added correctly to the database. WRONG PRICE. Expected: %f, Got: %f\n", newAlbum.Price, dbAlbum.Price)
			continue
		}

		// Delete the album from the database
		err = database.DeleteAlbum(id)

		// Check for errors
		if err != nil {
			t.Fatalf("Error deleting album from database: %s; Album: %v\n", err, newAlbum)
		}
	}
}
