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

func CreateBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := c.ShouldBindJSON(&books)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind JSON array: " + err.Error()})
		return
	}

	err = db.Create(&books).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot create Books: " + err.Error()})
		return
	}

	c.JSON(200, books)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditBook(c *gin.Context) {
	db := database.GetDatabase()

	var b models.Book

	err := c.ShouldBindJSON(&b)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&b).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, b)
}
