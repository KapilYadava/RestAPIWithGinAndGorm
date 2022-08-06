package controllers

import (
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Find books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": &books, "msg": "All records found !"})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &book, "msg": "Matching records found !"})
}

func CreateBook(c *gin.Context) {
	// validate input
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	//Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book, "msg": "Record created !"})
}

// delete all books record
func DeleteBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Delete(&books)
	c.JSON(http.StatusOK, gin.H{"data": &books, "msg": "All records deleted !"})
}

// delete a book record
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": &book, "msg": "Record deleted !"})
}

// Update a record
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// update book
	book = models.Book{Title: input.Title, Author: input.Author}
	models.DB.Update(&book)
	c.JSON(http.StatusOK, gin.H{"data": &book, "msg": "Record Updated !"})
}

//TODO
// Update all matching records
func UpdateBooks(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("Title = ?", c.Param("id")).Update(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !"})
		return
	}
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// update book
	book = models.Book{Title: input.Title, Author: input.Author}
	models.DB.Update(&book)
	c.JSON(http.StatusOK, gin.H{"data": &book, "msg": "All Matching Record Updated !"})
}
