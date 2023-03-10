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
		logger.FatalLogger.Fatalf("Error opening database: %s", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logger.FatalLogger.Fatalf("Error pinging database: %s", pingErr)
	}
	logger.InfoLogger.Println("Database connected!")
}

func CloseDB() {
	logger.Log.Println("Closing database connection...")
	db.Close()
}
