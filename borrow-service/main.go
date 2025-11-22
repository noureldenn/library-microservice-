package main

import (
	"borrow-service/config"
	"borrow-service/controllers"
	"borrow-service/models"
	"borrow-service/routes"
	"borrow-service/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Borrow{}) // or models.Borrow if your struct is in models

	r := gin.Default()

	bookClient := services.NewBookClient("http://localhost:3001")
	memberClient := services.NewMemberClient("http://localhost:4001")

	borrowController := controllers.NewBorrowController(bookClient, memberClient)

	routes.BorrowRoutes(r, borrowController)

	r.Run(":5001")
}
