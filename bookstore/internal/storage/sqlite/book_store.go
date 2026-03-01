package sqlite

import (
	"bookstore/internal/domain"
	"bookstore/internal/storage"
	"fmt"

	"gorm.io/gorm"
)

// SQLite book storage implementation
type BookStore struct {
	db *gorm.DB
}

// Create new SQLite BookStore
func NewBookStore(db *gorm.DB) storage.BookStorage {
	return &BookStore{db: db}
}

// Insert new book
func (s *BookStore) CreateBook(book *domain.Book) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}
	return s.db.Create(book).Error
}

// Fetch all books
func (s *BookStore) GetBooks() ([]domain.Book, error) {
	var books []domain.Book
	err := s.db.Find(&books).Error
	return books, err
}

// Fetch single book by ID
func (s *BookStore) GetBook(id uint) (*domain.Book, error) {
	var book domain.Book

	err := s.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// Delete book by ID
func (s *BookStore) DeleteBook(id uint) error {
	result := s.db.Delete(&domain.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
