package routes

import (
	"member-service/controllers"

	"github.com/gin-gonic/gin"
)

func MemberRoutes(r *gin.Engine) {
	r.GET("/members", controllers.GetAllMembers)
	r.POST("/members", controllers.CreateMember)
	r.GET("/members/:id", controllers.GetMemberByID)
	r.DELETE("/members/:id", controllers.DeleteMember)
}
