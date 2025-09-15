package controllers

import (
	"strconv"

	"github.com/Kovokar/BookApiGo/internal/database"
	"github.com/Kovokar/BookApiGo/internal/models"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	db := database.GetDatabase()

	var p []models.Book
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Id has to be a integer"})
		return
	}

	db := database.GetDatabase()

	err = db.First(&book, intId).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Book not found. Err: " + err.Error()})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind JSON: " + err.Error()})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot create Book: " + err.Error()})
		return
	}

	c.JSON(200, book)
}
