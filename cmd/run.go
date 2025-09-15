package cmd

import (
	"github.com/Kovokar/BookApiGo/internal/database"
	"github.com/Kovokar/BookApiGo/internal/server"
)

func Run() {

	database.StartDb()

	server := server.NewServer()
	server.Run()
}
