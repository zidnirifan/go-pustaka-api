package handler

import (
	"fmt"
	"log"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHanlder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (bookHandler *bookHandler) PostBook(c *gin.Context) {
	var body book.BodyPost

	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err.Error())

		validationErrors, isValidationError := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		if !isValidationError {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		for _, e := range validationErrors {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errorMessage,
			})
			return
		}
	}

	price, _ := body.Price.Int64()
	rate, _ := body.Rate.Int64()
	book := book.Book{Title: body.Title, Price: int(price), Description: body.Description, Rate: int(rate)}
	result, err := bookHandler.bookService.CreateBook(book)
	if err != nil {
		utils.HandleError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book successfully created",
		"data":    result,
	})
}

func (bookHandler *bookHandler) GetBooks(c *gin.Context) {
	// title := c.Query("title")
	books, err := bookHandler.bookService.GetBooks()
	if err != nil {
		utils.HandleError(err, c)
		return
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success get books",
		"data":    booksResponse,
	})
}

func (bookHandler *bookHandler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be number",
			"data":    err.Error(),
		})
		return
	}
	b, err := bookHandler.bookService.GetBookById(id)
	if err != nil {
		utils.HandleError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success get book",
		"data":    convertBookResponse(b),
	})
}

func (bookHandler *bookHandler) PutBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be number",
			"data":    err.Error(),
		})
	}

	var body book.BodyPut

	err = c.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err.Error())

		validationErrors, isValidationError := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		if !isValidationError {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		for _, e := range validationErrors {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errorMessage,
			})
			return
		}
	}

	price, _ := body.Price.Int64()
	rate, _ := body.Rate.Int64()
	book := book.Book{Title: body.Title, Price: int(price), Description: body.Description, Rate: int(rate)}
	result, err := bookHandler.bookService.UpdateBookById(id, book)
	if err != nil {
		utils.HandleError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book successfully updated",
		"data":    result,
	})
}

func (bookHandler *bookHandler) DeleteBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be number",
			"data":    err.Error(),
		})
	}

	id, err = bookHandler.bookService.DeleteBookById(id)
	if err != nil {
		utils.HandleError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "book successfully deleted",
		"data":    gin.H{"idDeleted": id},
	})
}

func convertBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rate:        b.Rate,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
