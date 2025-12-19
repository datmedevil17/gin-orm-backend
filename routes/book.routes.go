package routes

import (
	"gin-backend-orm/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	books := router.Group("/books")
	{
		books.POST("", controllers.CreateBook)
		books.GET("", controllers.GetBooks)
		books.GET("/:id", controllers.GetBook)
		books.PUT("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}
