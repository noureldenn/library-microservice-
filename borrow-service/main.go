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
	// 1️⃣ Connect to database
	config.ConnectDatabase()

	// 2️⃣ Auto-migrate Borrow model
	config.DB.AutoMigrate(&models.Borrow{}) // or models.Borrow if your struct is in models

	// 3️⃣ Create Gin router
	r := gin.Default()

	// 4️⃣ Initialize service clients
	bookClient := services.NewBookClient("http://localhost:3001")
	memberClient := services.NewMemberClient("http://localhost:4001")

	// 5️⃣ Create controller with injected clients
	borrowController := controllers.NewBorrowController(bookClient, memberClient)

	// 6️⃣ Register routes
	routes.BorrowRoutes(r, borrowController)

	// 7️⃣ Start server
	r.Run(":5001") // make sure port does not conflict with others
}
