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
