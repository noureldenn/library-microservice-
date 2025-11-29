package middleware

import (
	"auth-service/config"
	"auth-service/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("authoraization")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthenticated"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		config.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			log.Println("User not found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)

		c.Next()

	} else {
		c.JSON(401, gin.H{"error": "unauthenticated"})
		c.Abort()
		return
	}

	c.Next()
}
