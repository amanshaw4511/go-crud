package main

import (
	"crud/controller"
	"crud/repository"
	"crud/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := setup.ConnectDatabase()

	repo := repository.BookRepoImpl{Db: db}
	cont := controller.BookController{BookRepo: &repo}

	router.GET("/books", cont.GetAllBooks)
	router.POST("/books", cont.CreateBook)
	router.GET("/books/:id", cont.GetBookById)
	router.PUT("/books/:id", cont.UpdateBook)
	router.DELETE("/books/:id", cont.DeleteBook)

	router.Run()
}
