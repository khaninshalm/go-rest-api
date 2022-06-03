package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Name: "Hamlet", Author: "William Shakespeare"},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)

}

func postBook(context *gin.Context) {
	var new_book Book

	if err := context.BindJSON(&new_book); err != nil {
		return
	}

	books = append(books, new_book)
	context.IndentedJSON(http.StatusCreated, new_book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBook)
	router.Run()
}
