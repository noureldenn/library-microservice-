package main

import (
	"book-service/config"
	"book-service/models"
	"book-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Auto migrate models
	config.DB.AutoMigrate(&models.Author{}, &models.Category{}, &models.Publisher{}, &models.Book{})

	// Register routes
	routes.RegisterBookRoutes(r)

	// Start server
	r.Run(":3001") // Books service on port 3001
}
