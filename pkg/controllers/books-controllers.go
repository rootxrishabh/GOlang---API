package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/rootxrishabh/GOlang---API/pkg/models"
	"github.com/rootxrishabh/GOlang---API/pkg/utils"
	"fmt"
)

func GetBooks(c *gin.Context) {
	fmt.Println("ggg")
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	c.Header("Content-Type", "pkglication/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(res)
}

func GetBookById(c *gin.Context){
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, res)
}

func CreateBook(c *gin.Context){
	CreateBook := &models.Book{}
	utils.ParseBody(c.Request, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
}

func DeleteBook(c *gin.Context){
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	c.Header("Content-Type", "pkglication/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
}

func UpdateBook(c *gin.Context){
	var updateBook = &models.Book{}
	utils.ParseBody(c.Request, updateBook)
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	c.Writer.Header().Set("Content-Type", "pkglication/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
}