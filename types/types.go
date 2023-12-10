package types

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type Book struct {
	ID          uuid.UUID `db:"id"`
	AuthorID    uuid.UUID `db:"author_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type AuthorStore interface {
	Author(id uuid.UUID) (Author, error)
	AllAuthors() ([]Author, error)
	CreateAuthor(a *Author) error
	UpdateAuthor(a *Author) error
	DeleteAuthor(id *uuid.UUID) error
}

type BookStore interface {
	Book(id uuid.UUID) (Book, error)
	BooksByAuthor(authorID uuid.UUID) ([]Book, error)
	CreateBook(b *Book) error
	UpdateBook(b *Book) error
	DeleteBook(id uuid.UUID) error
}

type Store interface {
	AuthorStore
	BookStore
}
