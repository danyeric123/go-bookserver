package models

import (
	"github.com/danyeric123/go-bookserver/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title           string `gorm:""json:"title"`
	Author          uint
	Publisher       string `json:"publisher"`
	PublicationYear int64  `json:"publicationYear"`
}

type Author struct {
	gorm.Model
	FirstName string `json:"firstName`
	LastName  string `json:"lastName`
	DOB       string `json:"dob`
	Books     []Book `gorm:"foreignkey:BookID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
