package handler

import (
	"bookstore/internal/domain"
	"bookstore/internal/storage"
	"bookstore/internal/utils"
	"net/http"
)

type BookHandler struct {
	store storage.BookStorage
}

func NewBookHandler(store storage.BookStorage) *BookHandler {
	return &BookHandler{store: store}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.store.GetBooks()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(w, http.StatusOK, "Books fetched successfully", books)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Params parse error", err.Error())
	}

	book, _ := h.store.GetBook(id)
	utils.Success(w, http.StatusOK, "Book fetched successfully", book)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &domain.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid body", nil)
		return
	}

	if err := h.store.CreateBook(book); err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.Success(w, http.StatusCreated, "Book created", book)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Params parse error", err.Error())
	}

	book := h.store.DeleteBook(id)
	utils.Success(w, http.StatusOK, "Book deleted successfully", book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid 'id' parameter", err.Error())
		return
	}

	updateBook := &domain.Book{}
	if err := utils.ParseBody(r, updateBook); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	book, err := h.store.GetBook(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, "Book not found", err.Error())
		return
	}

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	if err := h.store.UpdateBook(book); err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to update book", err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "Book updated successfully", book)
}
