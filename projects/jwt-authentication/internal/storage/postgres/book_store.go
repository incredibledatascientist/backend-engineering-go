package postgres

import (
	"bookstore/internal/domain"
	"bookstore/internal/storage"
	"fmt"

	"gorm.io/gorm"
)

type BookStore struct {
	db *gorm.DB
}

func NewBookStore(db *gorm.DB) storage.BookStorage {
	return &BookStore{db: db}
}

func (s *BookStore) CreateBook(book *domain.Book) error {
	if book == nil {
		return fmt.Errorf("Book is nil")
	}

	return s.db.Create(book).Error
}

func (s *BookStore) GetBooks() ([]domain.Book, error) {
	var books []domain.Book
	err := s.db.Find(&books).Error
	return books, err
}

func (s *BookStore) GetBook(id uint) (*domain.Book, error) {
	var book domain.Book

	err := s.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

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

func (s *BookStore) UpdateBook(book *domain.Book) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}

	return s.db.Save(book).Error
}
