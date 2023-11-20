package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error")
	}
	log.Println("database connected successfully")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	bookV1 := router.Group("/v1/books")

	router.GET("/", handler.RootHanlder)
	bookV1.POST("/", bookHandler.PostBook)
	bookV1.GET("/", bookHandler.GetBooks)
	bookV1.GET("/:id", bookHandler.GetBookById)
	bookV1.PUT("/:id", bookHandler.PutBook)
	bookV1.DELETE("/:id", bookHandler.DeleteBookById)

	router.Run(":3000")
}
