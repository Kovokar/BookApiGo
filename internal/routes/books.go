package routes

import (
	"github.com/Kovokar/BookApiGo/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Função que registra as rotas relacionadas a "books"
func BooksRoutes(router *gin.RouterGroup) {
	books := router.Group("/books")
	{
		books.GET("/", controllers.ShowBooks)
		// books.POST("/", CreateBook)
		// books.GET("/:id", GetBookByID)
	}
}
