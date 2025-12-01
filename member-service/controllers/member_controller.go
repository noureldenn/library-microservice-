package controllers

import (
	"net/http"

	"member-service/config"
	"member-service/models"

	"github.com/gin-gonic/gin"
)

func GetAllMembers(c *gin.Context) {
	var members []models.Member
	config.DB.Find(&members)
	c.JSON(http.StatusOK, members)
}

func CreateMember(c *gin.Context) {
	var member models.Member

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, member)
}

func GetMemberByID(c *gin.Context) {
	id := c.Param("id")
	var member models.Member

	if err := config.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, member)
}

func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Member{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Member deleted"})
}
