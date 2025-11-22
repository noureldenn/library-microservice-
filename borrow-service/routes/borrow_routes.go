package routes

import (
	"borrow-service/controllers"
	"github.com/gin-gonic/gin"
)

func BorrowRoutes(r *gin.Engine, borrowController *controllers.BorrowController) {
	// POST /borrow
	r.POST("/borrow", borrowController.CreateBorrow)
}
