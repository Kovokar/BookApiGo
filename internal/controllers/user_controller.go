package controllers

import (
	"github.com/Kovokar/BookApiGo/internal/database"
	"github.com/Kovokar/BookApiGo/internal/models"
	"github.com/Kovokar/BookApiGo/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind JSON: " + err.Error()})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot create user: " + err.Error()})
		return
	}

	c.JSON(200, user)
}
