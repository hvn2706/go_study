package database

import (
	"database/sql"
	"gomock/logger"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Connectdb connects to the database
func Connectdb() {
	logger.Log.SetPrefix("[Connectdb] ")

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		logger.Log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logger.Log.Fatal(pingErr)
	}
	logger.Log.Println("Database connected!")
}

func CloseDB() {
	logger.Log.Println("Closing database connection...")
	db.Close()
}
