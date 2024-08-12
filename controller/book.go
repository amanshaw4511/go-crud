package controller

import (
	"crud/model"
	"crud/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookRepo repository.BookRepo
}

func (cont *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := cont.BookRepo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (cont *BookController) CreateBook(ctx *gin.Context) {

	var createBookRequest model.CreateBookRequest

	if err := ctx.ShouldBindJSON(&createBookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book, err := cont.BookRepo.Create(createBookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *BookController) GetBookById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var book model.Book
	book, err = cont.BookRepo.GetById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *BookController) UpdateBook(ctx *gin.Context) {
	var updateBookRequest model.UpdateBookRequest
	if err := ctx.ShouldBindJSON(&updateBookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book, err := cont.BookRepo.Update(uint(id), updateBookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book, err := cont.BookRepo.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, book)
}
