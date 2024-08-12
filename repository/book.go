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

type bookRepoImpl struct {
	db *gorm.DB
}

func BookRepoImpl(db *gorm.DB) BookRepo {
	return &bookRepoImpl{db: db}
}

func (repo *bookRepoImpl) GetAll() (books []model.Book, err error) {
	err = repo.db.Find(&books).Error
	return
}

func (repo *bookRepoImpl) GetById(id uint) (book model.Book, err error) {
	err = repo.db.First(&book, "id = ?", id).Error
	return
}

func (repo *bookRepoImpl) Create(createBookRequest model.CreateBookRequest) (book model.Book, err error) {
	book = model.Book{
		Title:  createBookRequest.Title,
		Author: createBookRequest.Author,
	}

	err = repo.db.Create(&book).Error
	return
}

func (repo *bookRepoImpl) Update(id uint, updateBookRequest model.UpdateBookRequest) (book model.Book, err error) {
	book, err = repo.GetById(id)
	if err != nil {
		return
	}

	err = repo.db.Model(&book).Updates(updateBookRequest).Error
	return
}

func (repo *bookRepoImpl) Delete(id uint) (book model.Book, err error) {
	book, err = repo.GetById(id)
	if err != nil {
		return
	}

	err = repo.db.Delete(&book).Error
	return
}
