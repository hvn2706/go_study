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
	}

	for _, newAlbum := range newAlbums {
		// Add the album to the database
		id, _ := service.AddAlbumToDB(newAlbum)

		if newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price < 0 {
			if id != service.ERROR_ID {
				t.Fatal("Returned ID is not ERROR_ID")
				continue
			} else {
				continue
			}
		}

		// Get the album from the database
		dbAlbum, err := database.GetAlbumByID(id)

		// Check for errors
		if err != nil {
			t.Fatal(err)
			continue
		}

		// Check if the album was added correctly
		if dbAlbum.Title != newAlbum.Title || dbAlbum.Artist != newAlbum.Artist || dbAlbum.Price != newAlbum.Price {
			t.Fatal("Album was not added correctly")
			continue
		}

		// Delete the album from the database
		err = database.DeleteAlbum(id)

		// Check for errors
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDeleteAlbumFromDB(t *testing.T) {
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
	}

	for _, newAlbum := range newAlbums {
		// Add the album to the database
		id, _ := service.AddAlbumToDB(newAlbum)

		if newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price < 0 {
			if id != service.ERROR_ID {
				t.Fatal("Returned ID is not ERROR_ID")
				continue
			} else {
				continue
			}
		}

		// Delete the album from the database
		err := service.DeleteAlbumFromDB(id)

		// Check for errors
		if err != nil {
			t.Fatal(err)
			continue
		}

		// Get the album from the database
		_, err = database.GetAlbumByID(id)

		// Check for errors
		if err == nil {
			t.Fatal("Album was not deleted")
			continue
		}
	}
}
