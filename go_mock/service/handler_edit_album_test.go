package service_test

import (
	"gomock/database"
	"gomock/logger"
	"gomock/service"
	"math"
	"testing"
)

func createRegularAlbums(t *testing.T) []int64 {
	// Create a new album
	newAlbums := []service.AlbumInfo{
		{
			Title:  "Test Title",
			Artist: "Test Artist",
			Price:  999999.99,
		},
	}

	returned_ids := make([]int64, 0)

	for _, newAlbum := range newAlbums {
		// Add the album to the database
		id, err := service.AddAlbumToDB(newAlbum)
		if err != nil {
			t.Fatalf("Error adding album to database: %v", err)
			return nil
		}
		returned_ids = append(returned_ids, id)
	}

	return returned_ids
}

func TestEditAlbumFromDB(t *testing.T) {
	// Init logger
	logger.Init(true, "edit_album")
	// Connect to the database
	database.Connectdb()
	defer database.CloseDB()

	ids := createRegularAlbums(t)

	if ids == nil {
		t.Fatal("Error creating regular albums")
	}

	editedAlbums := []service.AlbumStr{
		{
			ID:     ids[0],
			Title:  "Edited Title",
			Artist: "Edited Artist",
			Price:  "999999.99",
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 1",
			Artist: "Edited Artist 1",
			Price:  "0",
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 2",
			Artist: "Edited Artist 2",
			Price:  "-1",
		},
		{
			ID:     ids[0],
			Title:  "",
			Artist: "Edited Artist 3",
			Price:  "20",
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 4",
			Artist: "",
			Price:  "1282.3125",
		},
		{
			ID:     ids[0],
			Title:  "a",
			Artist: "b",
			Price:  "",
		},
		{
			ID:     ids[0],
			Title:  "",
			Artist: "",
			Price:  "",
		},
		{
			ID:     ids[0],
			Title:  "˙ˆ˙¨˙ˆ˙",
			Artist: "˙ˆ˙©ƒ†¥˙",
			Price:  "0.00000000001",
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 5",
			Artist: "Edited Artist 5",
			Price:  "0.00000000001",
		},
		{
			ID:     999999,
			Title:  "Edited Title 6",
			Artist: "Edited Artist 6",
			Price:  "0.00000000001",
		},
		{
			ID:     999999,
			Title:  "",
			Artist: "",
			Price:  "-1",
		},
		{
			ID:     999999,
			Title:  "",
			Artist: "Edited Artist 6",
			Price:  "0.00000000001",
		},
		{
			ID:     999999,
			Title:  "Edited Title 6",
			Artist: "",
			Price:  "0.00000000001",
		},
		{
			ID:     999999,
			Title:  "",
			Artist: "",
			Price:  "",
		},
		{
			ID:     -1,
			Title:  "",
			Artist: "",
			Price:  "",
		},
		{
			ID:     -1,
			Title:  "",
			Artist: "",
			Price:  "-1284",
		},
		{
			ID:     -1,
			Title:  "",
			Artist: "",
			Price:  "23876.12",
		},
	}

	expected_editedAlbums := []database.Album{
		{
			ID:     ids[0],
			Title:  "Edited Title",
			Artist: "Edited Artist",
			Price:  999999.99,
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 1",
			Artist: "Edited Artist 1",
			Price:  0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 1",
			Artist: "Edited Artist 3",
			Price:  20,
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 4",
			Artist: "Edited Artist 3",
			Price:  1282.3125,
		},
		{
			ID:     ids[0],
			Title:  "a",
			Artist: "b",
			Price:  1282.3125,
		},
		{
			ID:     ids[0],
			Title:  "a",
			Artist: "b",
			Price:  1282.3125,
		},
		{
			ID:     ids[0],
			Title:  "˙ˆ˙¨˙ˆ˙",
			Artist: "˙ˆ˙©ƒ†¥˙",
			Price:  0.00000000001,
		},
		{
			ID:     ids[0],
			Title:  "Edited Title 5",
			Artist: "Edited Artist 5",
			Price:  0.00000000001,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
		{
			ID:     0,
			Title:  "",
			Artist: "",
			Price:  0.0,
		},
	}

	for i, editedAlbum := range editedAlbums {
		// Edit the album in the database
		editedAlbum, _ := service.EditAlbumFromDB(editedAlbum)

		// Check if the album was edited correctly
		if editedAlbum.ID != expected_editedAlbums[i].ID {
			t.Fatalf("[Testcase: %d] Error editing album in database: expected id %v, got %v\n", i, expected_editedAlbums[i].ID, editedAlbum.ID)
		}

		if editedAlbum.Title != expected_editedAlbums[i].Title {
			t.Fatalf("[Testcase: %d] Error editing album in database: expected title %v, got %v\n", i, expected_editedAlbums[i].Title, editedAlbum.Title)
		}

		if editedAlbum.Artist != expected_editedAlbums[i].Artist {
			t.Fatalf("[Testcase: %d] Error editing album in database: expected artist %v, got %v\n", i, expected_editedAlbums[i].Artist, editedAlbum.Artist)
		}

		if math.Abs(editedAlbum.Price-expected_editedAlbums[i].Price) > service.ACCEPTED_PRICE_ERROR {
			t.Fatalf("[Testcase: %d] Error editing album in database: expected price %f, got %f, sub: %f\n", i, expected_editedAlbums[i].Price, editedAlbum.Price, math.Abs(editedAlbum.Price-expected_editedAlbums[i].Price))
		}
	}

	// Delete the albums from the database
	for _, id := range ids {
		err := service.DeleteAlbumFromDB(id)
		if err != nil {
			t.Fatalf("Error deleting album from database: %v\n", err)
		}
	}
}
