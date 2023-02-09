package main

import (
	"gomock/database"
	"gomock/service"
)

func main() {
	defer database.CloseDB()
	database.Connectdb()
	service.RunServer()
}
