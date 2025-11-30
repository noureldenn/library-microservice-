package controllers

import (
	"book-service/config"
	"book-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Preload("Author").Preload("Category").Preload("Publisher").Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.Preload("Author").Preload("Category").Preload("Publisher").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result models.Book
	if err := config.DB.
		Preload("Author").
		Preload("Category").
		Preload("Publisher").
		First(&result, book.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book created successfully",
		"book":    result,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
func DecreaseAvailability(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no available copies"})
		return
	}

	config.DB.Model(&book).Update("available_copies", book.AvailableCopies-1)

	c.JSON(http.StatusOK, gin.H{"message": "available copies decreased"})
}
