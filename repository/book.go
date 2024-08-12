package repository

import (
	"crud/model"

	"gorm.io/gorm"
)

type BookRepo interface {
	GetAll() ([]model.Book, error)
	GetById(uint) (model.Book, error)
	Create(model.CreateBookRequest) (model.Book, error)
	Update(uint, model.UpdateBookRequest) (model.Book, error)
	Delete(uint) (model.Book, error)
}

type BookRepoImpl struct {
	Db *gorm.DB
}

func (repo *BookRepoImpl) GetAll() (books []model.Book, err error) {
	err = repo.Db.Find(&books).Error
	return
}

func (repo *BookRepoImpl) GetById(id uint) (book model.Book, err error) {
	err = repo.Db.First(&book, "id = ?", id).Error
	return
}

func (repo *BookRepoImpl) Create(createBookRequest model.CreateBookRequest) (book model.Book, err error) {
	book = model.Book{
		Title:  createBookRequest.Title,
		Author: createBookRequest.Author,
	}

	err = repo.Db.Create(&book).Error
	return
}

func (repo *BookRepoImpl) Update(id uint, updateBookRequest model.UpdateBookRequest) (book model.Book, err error) {
	book, err = repo.GetById(id)
	if err != nil {
		return
	}

	err = repo.Db.Model(&book).Updates(updateBookRequest).Error
	return
}

func (repo *BookRepoImpl) Delete(id uint) (book model.Book, err error) {
	book, err = repo.GetById(id)
	if err != nil {
		return
	}

	err = repo.Db.Delete(&book).Error
	return
}
