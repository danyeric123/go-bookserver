package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/danyeric123/go-bookserver/pkg/models"
	"github.com/danyeric123/go-bookserver/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func retrieveBookId(r *http.Request) int64{
	bookId := mux.Vars(r)["bookId"]
	idNum, err := strconv.ParseInt(bookId, 0, 0) 
	if err != nil {
		fmt.Println(err)
	}
	return idNum
}

func Index(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to Go-Bookserver"))
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	utils.SendJsonResponse(w, res)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	bookId := retrieveBookId(r)
	book, _ := models.GetBook(bookId)
	res, _ := json.Marshal(book)
	utils.SendJsonResponse(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	utils.SendJsonResponse(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	bookId := retrieveBookId(r)
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	utils.SendJsonResponse(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	bookId := retrieveBookId(r)
	bookDetails, db:=models.GetBook(bookId)
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
	utils.SendJsonResponse(w, res)
}