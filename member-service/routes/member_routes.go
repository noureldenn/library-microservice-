package routes

import (
	"member-service/controllers"

	"member-service/middlewares"

	"github.com/gin-gonic/gin"
)

func MemberRoutes(r *gin.Engine) {
	memberroutes := r.Group("/members")

	{
		memberroutes.GET("/", controllers.GetAllMembers)
		memberroutes.GET("/:id", controllers.GetMemberByID)

		memberroutes.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
		{

			memberroutes.POST("/", controllers.CreateMember)
			memberroutes.DELETE("/:id", controllers.DeleteMember)
		}
	}
}
