package main

import (
	"github.com/gin-gonic/gin"
	"member-service/config"
	"member-service/models"
	"member-service/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Member{})

	routes.MemberRoutes(r)

	r.Run(":4001")
}
