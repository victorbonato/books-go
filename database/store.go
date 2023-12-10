package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewStore(databaseURL string) (*Store, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		AuthorStore: &AuthorStore{DB: db},
		BookStore:   &BookStore{DB: db},
	}, nil
}

type Store struct {
	*AuthorStore
	*BookStore
}
