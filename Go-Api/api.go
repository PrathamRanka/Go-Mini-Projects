package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "a", Author: "b", Quantity: 1},
	{ID: "2", Title: "c", Author: "d", Quantity: 2},
	{ID: "3", Title: "e", Author: "f", Quantity: 3},
}

//get endpoint
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

//post endpoint
func createBooks(c *gin.Context){
	var newbook book

	if err := c.BindJSON(&newbook); //bindjson used to serve back error
	err != nil {
		return
	}

	books = append(books, newbook)
	c.IndentedJSON(http.StatusCreated, newbook)
}

//fetching 
// first is for input parameters, second is for return type
 func getBookbyId(id string)(*book, error){
	  for i, b := range books {
		if b.ID = id {
			return &books[i], nil
		}
	  }
	  return  nil, errors.New("book not found")
 }

 func bookbyId(c *gin.Context) {
	id := c.Params("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON((http.StatusOK), book)
 }

func main() {
	router := gin.Default()
	// println("Server is started at 8080")
	router.GET("/books" , getBooks)
	router.GET("/books/:id" , bookbyId)
	router.POST("/books" , createBooks)
	router.Run("localhost:8080")
}