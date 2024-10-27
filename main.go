package main

import (
	"log"
	"os"

	"github.com/felipeselau/go-books-api/database"
	"github.com/felipeselau/go-books-api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	str := os.Getenv("DATABASE_URL")
	if str == "" {
		log.Fatal("Error loading DATABASE_URL")
	}

	database.StartDB(str)

	app := server.NewServer()

	server.Run(&app)
}
