package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mayankdev-oss/go-bookstore-sql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `gorm:"" json:"Author"`
	Publication string `gorm:"" json:"Publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var book []Book
	db.Find(&book)
	return book
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	res := db.Where("id = ?", Id).First(&getBook)
	return &getBook, res
}
