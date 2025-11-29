package controllers

import (
	"book-service/config"
	"book-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	config.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	var input models.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&author).Updates(input)
	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := config.DB.Delete(&author, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted success"})
}
