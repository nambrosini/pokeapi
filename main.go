package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nambrosini/pokeapi/controllers"
	"github.com/nambrosini/pokeapi/models"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()

	models.ConnectDatabase()
	models.ConnectRedis()

	r.GET("/books", controllers.GetAllBooks)
	r.POST("/book", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	log.Println("Listening on port " + port)
	log.Fatal(r.Run(":" + port))
}
