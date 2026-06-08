package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mayankdev-oss/go-bookstore-sql/pkg/models"
	"github.com/mayankdev-oss/go-bookstore-sql/pkg/utils"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	newbooks := models.GetAllBooks()
	//data, _ := json.Marshal(newbooks)
	//w.Write(data)
	err := json.NewEncoder(w).Encode(newbooks)
	if err != nil {
		log.Print("Error in encoding the data into JSON\n", http.StatusBadRequest)
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := &models.Book{}
	w.WriteHeader(http.StatusOK)
	utils.ParseBody(r, CreateBook)
	// data, _ := json.Marshal(CreateBook)
	// json.NewEncoder(w).Encode(data)
	// w.WriteHeader(http.StatusOK)
	b := models.CreateBook(CreateBook)
	err := json.NewEncoder(w).Encode(b)
	if err != nil {
		log.Print("Error in encoding the data into JSON\n", http.StatusBadRequest)

	}

}

func GetBookbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	id, _ := strconv.ParseInt(bookId, 10, 64)
	book, _ := models.GetBookById(id)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	update := &models.Book{}
	utils.ParseBody(r, update)
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	id, _ := strconv.ParseInt(bookId, 10, 64)
	bookdetails, db := models.GetBookById(id)
	if update.Name != "" {
		bookdetails.Name = update.Name
	}

	if update.Author != "" {
		bookdetails.Author = update.Author
	}

	if update.Publication != "" {
		bookdetails.Publication = update.Publication
	}

	db.Save(&update)
	json.NewEncoder(w).Encode(bookdetails)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookid"]
	id, _ := strconv.ParseInt(bookId, 10, 64)
	book := models.DeleteBook(id)
	json.NewEncoder(w).Encode(book)

}
