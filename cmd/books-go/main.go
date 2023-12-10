package main

import (
	"log"
	"os"

	"github.com/victorbonato/books-go/database"
	"github.com/victorbonato/books-go/web"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db_url := os.Getenv("DATABASE_URL")

	store, err := database.NewStore(db_url)
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	if err := h.Run(":3000"); err != nil {
		return
	}
}
