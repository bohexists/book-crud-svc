package service

import (
	"github.com/bohexists/book-crud-svc/internal/domain"
)

// BookService service for books
type BookService struct {
	Repo BookServiceInterface
}

// NewBookService creates a new BookService
func NewBookService(repo BookServiceInterface) *BookService {
	return &BookService{
		Repo: repo,
	}
}

type BookServiceInterface interface {
	GetBooks() ([]domain.Book, error)
	GetBook(id int) (domain.Book, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id int, book domain.Book) error
	DeleteBook(id int) error
}

// GetBooks returns all books
func (s *BookService) GetBooks() ([]domain.Book, error) {
	return s.Repo.GetBooks()
}

// GetBook returns a single book
func (s *BookService) GetBook(id int) (domain.Book, error) {
	return s.Repo.GetBook(id)
}

// CreateBook creates a new book
func (s *BookService) CreateBook(book domain.Book) (domain.Book, error) {
	return s.Repo.CreateBook(book)
}

// UpdateBook updates an existing book
func (s *BookService) UpdateBook(id int, book domain.Book) error {
	return s.Repo.UpdateBook(id, book)
}

// DeleteBook deletes a book
func (s *BookService) DeleteBook(id int) error {
	return s.Repo.DeleteBook(id)
}
