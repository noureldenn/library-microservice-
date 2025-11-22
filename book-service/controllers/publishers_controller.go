package controllers

import (
	"book-service/config"
	"book-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new publisher
func CreatePublisher(c *gin.Context) {
	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&publisher)
	c.JSON(http.StatusCreated, publisher)
}

// Get all publishers
func GetPublishers(c *gin.Context) {
	var publishers []models.Publisher
	config.DB.Find(&publishers)
	c.JSON(http.StatusOK, publishers)
}

// Get publisher by ID
func GetPublisherByID(c *gin.Context) {
	id := c.Param("id")
	var publisher models.Publisher

	if err := config.DB.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}
	c.JSON(http.StatusOK, publisher)
}

// Update publisher
func UpdatePublisher(c *gin.Context) {
	id := c.Param("id")
	var publisher models.Publisher

	if err := config.DB.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}

	var input models.Publisher
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&publisher).Updates(input)
	c.JSON(http.StatusOK, publisher)
}

// Delete publisher
func DeletePublisher(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Publisher{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Publisher deleted successfully"})
}
