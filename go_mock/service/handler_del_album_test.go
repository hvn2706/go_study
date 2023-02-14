package service_test

import (
	"gomock/database"
	"gomock/service"
	"testing"
)

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
		{
			Title:  "",
			Artist: "",
			Price:  -1,
		},
		{
			Title:  "˙ˆ˙¨˙ˆ˙",
			Artist: "˙ˆ˙©ƒ†¥˙",
			Price:  0.00000000001,
		},
	}

	var deleting_ids []int64 = make([]int64, 0)

	for i, newAlbum := range newAlbums {
		// Add the album to the database
		id, _ := service.AddAlbumToDB(newAlbum)

		if newAlbum.Title == service.NIL_TITLE || newAlbum.Artist == service.NIL_ARTIST || newAlbum.Price < service.MIN_PRICE_NUM {
			if id != service.ERROR_ID {
				t.Fatalf("[Testcase: %d] Returned ID is not ERROR_ID for album: %v\n", i, newAlbum)
				continue
			} else {
				continue
			}
		}

		deleting_ids = append(deleting_ids, id)
	}

	deleting_ids = append(deleting_ids, 0)
	deleting_ids = append(deleting_ids, -1)
	deleting_ids = append(deleting_ids, 8882)

	for i, id := range deleting_ids {
		// Delete the album from the database
		service.DeleteAlbumFromDB(id)

		// Get the album from the database
		_, err := database.GetAlbumByID(id)

		// Check for errors
		if err == nil {
			t.Fatalf("[Testcase: %d] Album was not deleted from the database: %s; ID: %d\n", i, err, id)
			continue
		}
	}
}
