package database

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/victorbonato/books-go/types"
)

type BookStore struct {
	*sqlx.DB
}

func (s *BookStore) Book(id uuid.UUID) (types.Book, error) {
	var b types.Book
	if err := s.Get(&b, `SELECT * FROM books WHERE id = $1`, id); err != nil {
		return types.Book{}, fmt.Errorf("error getting book: %w", err)
	}
	return b, nil
}

func (s *BookStore) BooksByAuthor(authorID uuid.UUID) ([]types.Book, error) {
	var bb []types.Book
	if err := s.Get(&bb, `SELECT * FROM books WHERE author_id = $1`, authorID); err != nil {
		return []types.Book{}, fmt.Errorf("error getting books by id: %w", err)
	}
	return bb, nil
}

func (s *BookStore) CreateBook(b *types.Book) error {
	if err := s.Get(b, `INSERT INTO books VALUES ($1, $2, $3, $4) RETURNING *`,
		b.ID,
		b.AuthorID,
		b.Title,
		b.Description); err != nil {
		return fmt.Errorf("error creating new book: %w", err)
	}
	return nil
}

func (s *BookStore) UpdateBook(b *types.Book) error {
	if err := s.Get(b, `UPDATE books SET author_id = $1, title = $2, description = $3 WHERE id = $4 RETURNING *`,
		b.AuthorID,
		b.Title,
		b.Description,
		b.ID); err != nil {
		return fmt.Errorf("error updating book: %w", err)
	}
	return nil
}

func (s *BookStore) DeleteBook(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM books WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting book: %w", err)
	}
	return nil
}
