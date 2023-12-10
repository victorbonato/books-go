package main

import (
	"log"

	"github.com/victorbonato/books-go/database"
	"github.com/victorbonato/books-go/web"
)

func main() {
	db_url := "postgres://postgres:password@localhost/postgres?sslmode=disable"
	store, err := database.NewStore(db_url)
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	if err := h.Run(":3000"); err != nil {
		return
	}
}
