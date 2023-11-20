package book

import (
	"errors"
	"net/http"
	"pustaka-api/utils"
)

type Service interface {
	CreateBook(data Book) (Book, error)
	GetBooks() ([]Book, error)
	GetBookById(id int) (Book, error)
	UpdateBookById(id int, dataUpdated Book) (Book, error)
	DeleteBookById(id int) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r *repository) *service {
	return &service{r}
}

func (s *service) CreateBook(data Book) (Book, error) {
	return s.repository.CreateBook(data)
}

func (s *service) GetBooks() ([]Book, error) {
	return s.repository.GetBooks()
}

func (s *service) GetBookById(id int) (Book, error) {
	book, err := s.repository.GetBookById(id)
	if err != nil {
		return Book{}, &utils.RequestError{StatusCode: http.StatusNotFound, Err: errors.New("book not found")}
	}
	return book, nil
}

func (s *service) UpdateBookById(id int, dataUpdated Book) (Book, error) {
	book, err := s.repository.GetBookById(id)
	if err != nil {
		return Book{}, &utils.RequestError{StatusCode: http.StatusNotFound, Err: errors.New("book not found")}
	}

	if dataUpdated.Title != "" {
		book.Title = dataUpdated.Title
	}
	if dataUpdated.Description != "" {
		book.Description = dataUpdated.Description
	}
	if dataUpdated.Price != 0 {
		book.Price = dataUpdated.Price
	}
	if dataUpdated.Rate != 0 {
		book.Rate = dataUpdated.Rate
	}

	return s.repository.UpdateBookById(book)
}

func (s *service) DeleteBookById(id int) (int, error) {
	_, err := s.repository.GetBookById(id)
	if err != nil {
		return id, &utils.RequestError{StatusCode: http.StatusNotFound, Err: errors.New("book not found")}
	}

	return s.repository.DeleteBookById(id)
}
