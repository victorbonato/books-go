package database

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/victorbonato/books-go/types"
)

type AuthorStore struct {
	*sqlx.DB
}

func (s *AuthorStore) Author(id uuid.UUID) (types.Author, error) {
	var a types.Author
	if err := s.Get(&a, `SELECT * FROM authors WHERE id = $1`, id); err != nil {
		return types.Author{}, fmt.Errorf("error getting author: %w", err)
	}
	return a, nil
}

func (s *AuthorStore) AllAuthors() ([]types.Author, error) {
	var aa []types.Author
	if err := s.Select(&aa, `SELECT * FROM authors`); err != nil {
		return []types.Author{}, fmt.Errorf("error getting all authors: %w", err)
	}
	return aa, nil
}

func (s *AuthorStore) CreateAuthor(a *types.Author) error {
	if err := s.Get(a, `INSERT INTO authors VALUES ($1, $2) RETURNING *`,
		a.ID,
		a.Name); err != nil {
		return fmt.Errorf("error creating new author %w", err)
	}
	return nil
}

func (s *AuthorStore) UpdateAuthor(a *types.Author) error {
	if err := s.Get(a, `UPDATE authors SET name = $1 WHERE id = $2 RETURNING *`,
		a.Name,
		a.ID); err != nil {
		return fmt.Errorf("error updating author: %w", err)
	}
	return nil
}

func (s *AuthorStore) DeleteAuthor(id *uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM authors WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting author: %w", err)
	}
	return nil
}
