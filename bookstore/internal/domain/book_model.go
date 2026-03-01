package domain

import (
	"bookstore/internal/config"
	"fmt"

	"gorm.io/gorm"
	// "jwt-auth/internal/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model       // Gives:- ID(Primary Key), CreatedAt, UpdatedAt, DeletedAt
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type CreateBookRequest struct {
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// config

func init() {
	err := config.Connect()
	if err != nil {
		fmt.Println("Config err book model:", err.Error())
		return
	}

	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// func (b *Book) CreateBook() *Book {
// func (b *Book) CreateBook() (*Book, error) {
func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// func (b *Book) GetBooks() []*Book {
func GetBooks() []*Book {
	// db.NewRecord(b)
	books := []*Book{}
	db.Find(&books)
	return books
}

// func (b *Book) GetBook(id int) (*Book, *gorm.DB) {
func GetBook(id int) (*Book, *gorm.DB) {
	book := &Book{}
	db.Where("ID=?", id).Find(book)
	return book, db
}

// func (b *Book) UpdateBook(id int, b *Book) (*Book, *gorm.DB) {
// 	book := &Book{}
// 	db.Where("ID=?", id).Find(book)
// 	return book, db
// }

// func (b *Book) DeleteBook(id int) Book {
func DeleteBook(id int) Book {
	book := Book{}
	db.Where("ID=?", id).Find(&book)
	return book
}
