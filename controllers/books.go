package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nambrosini/pokeapi/models"
	"net/http"
)

// GetAllBooks GET /books
// Get all books
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	book.Create()

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	book, err := models.GetBook(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /book/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	book, err := models.GetBook(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.UpdateBook(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /book/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	book, err := models.GetBook(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erorr": "Record not found"})
		return
	}

	book.DeleteBook()
	c.JSON(http.StatusOK, gin.H{"data": book})
}
