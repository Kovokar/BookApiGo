package routes

import (
	"github.com/Kovokar/BookApiGo/internal/controllers"
	"github.com/Kovokar/BookApiGo/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Função que registra as rotas relacionadas a "books"
func BooksRoutes(router *gin.RouterGroup) {
	books := router.Group("/books", middleware.Auth())
	{
		books.GET("/", controllers.GetBook)
		books.GET("/:id", controllers.GetBookById)
		books.POST("/", controllers.CreateBook)
		books.POST("/bulk", controllers.CreateBooks)
		books.PUT("/:id", controllers.EditBook)
		books.DELETE("/:id", controllers.DeleteBook)

		// books.POST("/", CreateBook)
		// books.GET("/:id", GetBookByID)
	}
}
