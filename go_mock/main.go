package main

import (
	"gomock/database"
	"gomock/logger"
	"gomock/service"
)

func main() {
	defer database.CloseDB()
	logger.Init(false, "")
	database.Connectdb()
	service.RunServer()
}
