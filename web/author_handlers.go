package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorbonato/books-go/types"
)

func (h *Handler) GetAuthors(c *gin.Context) {
	aa, err := h.store.AllAuthors()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": aa})
}

func (h *Handler) GetAuthorByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.store.Author(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": a})
}

type AuthorInput struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) CreateAuthor(c *gin.Context) {
	var input AuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAuthor := &types.Author{
		ID:   uuid.New(),
		Name: input.Name,
	}
	if err := h.store.CreateAuthor(newAuthor); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAuthor})
}

func (h *Handler) UpdateAuthor(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input AuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAuthor := &types.Author{
		ID:   id,
		Name: input.Name,
	}
	if err := h.store.UpdateAuthor(updatedAuthor); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": updatedAuthor})
}

func (h *Handler) DeleteAuthor(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.DeleteAuthor(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"data": true})
}
