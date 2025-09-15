package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		BooksRoutes(main)
	}

	return router
}
