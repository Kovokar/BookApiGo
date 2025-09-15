package controllers

import (
	"github.com/Kovokar/BookApiGo/internal/database"
	"github.com/Kovokar/BookApiGo/internal/models"
	"github.com/Kovokar/BookApiGo/internal/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind JSON: " + err.Error()})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{"error": "Cannot find user: "})
		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.NewJWTService().GerenateToken(uint(user.ID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
