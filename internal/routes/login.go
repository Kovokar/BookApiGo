package routes

import (
	"github.com/Kovokar/BookApiGo/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Função que registra as rotas relacionadas a "books"
func LoginRoutes(router *gin.RouterGroup) {
	user := router.Group("/login")
	{
		user.POST("/", controllers.Login)
	}
}
