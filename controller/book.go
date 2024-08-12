package controller

import (
	"crud/model"
	"crud/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	repo repository.BookRepo
}

func BookController(repo repository.BookRepo) *bookController {
	return &bookController{repo: repo}
}

func (cont *bookController) GetAllBooks(ctx *gin.Context) {
	books, err := cont.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (cont *bookController) CreateBook(ctx *gin.Context) {

	var createBookRequest model.CreateBookRequest

	if err := ctx.ShouldBindJSON(&createBookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book, err := cont.repo.Create(createBookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *bookController) GetBookById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var book model.Book
	book, err = cont.repo.GetById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *bookController) UpdateBook(ctx *gin.Context) {
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

	book, err := cont.repo.Update(uint(id), updateBookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *bookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book, err := cont.repo.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, book)
}
