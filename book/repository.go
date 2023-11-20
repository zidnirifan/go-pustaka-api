package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateBook(data Book) (Book, error)
	GetBooks() ([]Book, error)
	GetBookById(id int) (Book, error)
	UpdateBookById(dataUpdated Book) (Book, error)
	DeleteBookById(id int) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(data Book) (Book, error) {
	result := r.db.Create(&data)
	return data, result.Error
}

func (r *repository) GetBooks() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) GetBookById(id int) (Book, error) {
	var book Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *repository) UpdateBookById(dataUpdated Book) (Book, error) {
	err := r.db.Save(dataUpdated).Error
	return dataUpdated, err
}

func (r *repository) DeleteBookById(id int) (int, error) {
	err := r.db.Delete(&Book{}, id).Error
	return id, err
}
