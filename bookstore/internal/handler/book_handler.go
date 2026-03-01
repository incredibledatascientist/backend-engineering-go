package handler

import (
	"bookstore/internal/domain"
	"bookstore/internal/utils"
	"net/http"
)

var NewBook domain.Book

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

// func (b *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := domain.GetBooks()
	// res, err := json.Marshal(books)
	// if err != nil {
	// 	utils.Error(w, http.StatusInternalServerError, "Json marshal failed", err.Error())
	// }
	utils.Success(w, http.StatusOK, "Books fetched successfully", books)

}

// func (b *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Params parse error", err.Error())
	}

	book, _ := domain.GetBook(id)
	utils.Success(w, http.StatusOK, "Book fetched successfully", book)
}

// func (b *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &domain.Book{}
	// But realtime applications uses schemas they dont want to exposed all field in the fronted.
	// Mostly use Request/Response models
	if err := utils.ParseBody(r, newBook); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	newBook.ID = 0 // prevent manual ID injection

	book, err := newBook.CreateBook()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.Success(w, http.StatusCreated, "Book created successfully", book)
}

// func (b *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Params parse error", err.Error())
	}

	book := domain.DeleteBook(id)
	utils.Success(w, http.StatusOK, "Book deleted successfully", book)
}

// func (b *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Params parse error", err.Error())
	}

	updateBook := &domain.Book{}
	utils.ParseBody(r, updateBook)

	book, db := domain.GetBook(id)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)
	utils.Success(w, http.StatusOK, "Book updated successfully", book)
}
