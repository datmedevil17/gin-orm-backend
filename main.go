package main

import (
	"log"

	"gin-backend-orm/database"
	"gin-backend-orm/models"
	"gin-backend-orm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	err := database.DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	routes.BookRoutes(r)
	r.Run(":8080")
}
