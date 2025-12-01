package routes

import (
	"book-service/controllers"
	"book-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", controllers.GetAllBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.PUT("/:id/decrease", controllers.DecreaseAvailability)

		bookRoutes.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
		{
			bookRoutes.POST("/", controllers.AddBook)
			bookRoutes.PUT("/:id", controllers.UpdateBook)
			bookRoutes.DELETE("/:id", controllers.DeleteBook)

		}

	}

	authorRoutes := r.Group("/authors")
	{
		authorRoutes.GET("/", controllers.GetAuthors)
		authorRoutes.GET("/:id", controllers.GetAuthorByID)

		authorRoutes.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
		{
			authorRoutes.POST("/", controllers.CreateAuthor)
			authorRoutes.PUT("/:id", controllers.UpdateAuthor)
			authorRoutes.DELETE("/:id", controllers.DeleteAuthor)

		}
	}

	categoryRoutes := r.Group("/categories")
	{
		categoryRoutes.GET("/", controllers.GetCategories)
		categoryRoutes.GET("/:id", controllers.GetCategoryByID)

		categoryRoutes.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
		{
			categoryRoutes.POST("/", controllers.CreateCategory)
			categoryRoutes.PUT("/:id", controllers.UpdateCategory)
			categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
		}
	}

	publisherRoutes := r.Group("/publishers")
	{
		publisherRoutes.GET("/", controllers.GetPublishers)
		publisherRoutes.GET("/:id", controllers.GetPublisherByID)

		publisherRoutes.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
		{
			publisherRoutes.POST("/", controllers.CreatePublisher)
			publisherRoutes.PUT("/:id", controllers.UpdatePublisher)
			publisherRoutes.DELETE("/:id", controllers.DeletePublisher)
		}
	}
}
