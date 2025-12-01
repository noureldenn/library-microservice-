package routes

import (
	"borrow-service/controllers"
	"borrow-service/middlewares"

	"github.com/gin-gonic/gin"
)

func BorrowRoutes(r *gin.Engine, borrowController *controllers.BorrowController) {
	// POST /borrow
	r.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
	r.POST("/borrow", borrowController.CreateBorrow)
}
