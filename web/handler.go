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
	books_routes := h.Group("/books")
	{
		books_routes.GET("/", h.GetAllBooksInfo)
		books_routes.POST("/", h.CreateBook)
		books_routes.GET("/:id", h.GetBookByID)
		books_routes.GET("/author/:id", h.GetBooksByAuthor)
		books_routes.PATCH("/:id", h.UpdateBook)
		books_routes.DELETE(":id", h.DeleteBook)
	}

	return h
}

type Handler struct {
	*gin.Engine

	store types.Store
}
