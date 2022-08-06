package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.CreateDatabase()
	defer models.DB.Close()
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.DELETE("/books", controllers.DeleteBooks)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.PUT("/books/:id", controllers.UpdateBook)

	r.Run()
}
