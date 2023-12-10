package web

import (
	"net/http"

	"github.com/victorbonato/books-go/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetAllBooksInfo(c *gin.Context) {
	bbi, err := h.store.BooksWithAuthors()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": bbi})
}

func (h *Handler) GetBookByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b, err := h.store.Book(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": b})
}

func (h *Handler) GetBooksByAuthor(c *gin.Context) {
	authorID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bb, err := h.store.BooksByAuthor(authorID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": bb})
}

type BookInput struct {
	AuthorID    uuid.UUID `json:"author_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

func (h *Handler) CreateBook(c *gin.Context) {
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook := &types.Book{
		ID:          uuid.New(),
		AuthorID:    input.AuthorID,
		Title:       input.Title,
		Description: input.Description,
	}
	if err := h.store.CreateBook(newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newBook})
}

func (h *Handler) UpdateBook(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook := &types.Book{
		ID:          id,
		AuthorID:    input.AuthorID,
		Title:       input.Title,
		Description: input.Description,
	}
	if err := h.store.UpdateBook(updatedBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": updatedBook})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.DeleteBook(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"data": true})
}
