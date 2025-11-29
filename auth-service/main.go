package main

import (
	"auth-service/config"
	"auth-service/controllers"
	"auth-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.Run(":3000")
}
