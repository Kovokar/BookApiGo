package cmd

import "github.com/Kovokar/BookApiGo/internal/server"

func Run() {
	server := server.NewServer()
	server.Run()
}
