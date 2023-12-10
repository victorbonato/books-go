package web

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbonato/books-go/types"
)

func NewHandler(store types.Store) *Handler {
	h := &Handler{
		Engine: gin.Default(),
		store:  store,
	}

	authors_routes := h.Group("/authors")
	{
		authors_routes.GET("/", h.GetAuthors)
		authors_routes.POST("/", h.CreateAuthor)
		authors_routes.GET("/:id", h.GetAuthorByID)
		authors_routes.PATCH("/:id", h.UpdateAuthor)
		authors_routes.DELETE("/:id", h.DeleteAuthor)
	}

	return h
}

type Handler struct {
	*gin.Engine

	store types.Store
}
