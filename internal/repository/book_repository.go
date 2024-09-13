package repository

import (
	"database/sql"
	"github.com/bohexists/book-crud-svc/internal/domain"
	"github.com/sirupsen/logrus"
)

// BookRepository структура для работы с базой данных по книгам
type BookRepository struct {
	DB  *sql.DB
	Log *logrus.Logger
}

// NewBookRepository создает новый BookRepository
func NewBookRepository(db *sql.DB, log *logrus.Logger) BookRepositoryInterface {
	return &BookRepository{
		DB:  db,
		Log: log,
	}
}

// BookRepositoryInterface интерфейс для работы с базой данных по книгам
type BookRepositoryInterface interface {
	GetBooks() ([]domain.Book, error)
	GetBook(id int) (domain.Book, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id int, book domain.Book) error
	DeleteBook(id int) error
}

// GetBooks извлекает все книги из базы данных
func (r *BookRepository) GetBooks() ([]domain.Book, error) {
	rows, err := r.DB.Query("SELECT id, title, description, author, published, price FROM books")
	if err != nil {
		r.Log.WithError(err).Error("Failed to fetch books")
		return nil, err
	}

	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		r.Log.WithError(err).Error("Failed to fetch books")
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Published, &book.Price); err != nil {
			r.Log.WithError(err).Error("Failed to scan book")
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// GetBook извлекает одну книгу по ID
func (r *BookRepository) GetBook(id int) (domain.Book, error) {
	var book domain.Book
	err := r.DB.QueryRow("SELECT id, title, description, author, published, price FROM books WHERE id = $1", id).Scan(
		&book.ID, &book.Title, &book.Description, &book.Author, &book.Published, &book.Price)
	if err != nil {
		r.Log.WithError(err).Errorf("Failed to fetch book with ID: %d", id)
		return domain.Book{}, err
	}
	return book, nil
}

// CreateBook добавляет новую книгу в базу данных
func (r *BookRepository) CreateBook(book domain.Book) (domain.Book, error) {
	err := r.DB.QueryRow("INSERT INTO books (title, description, author, published, price) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		book.Title, book.Description, book.Author, book.Published, book.Price).Scan(&book.ID)
	if err != nil {
		r.Log.WithError(err).Error("Failed to create book")
		return domain.Book{}, err
	}
	return book, nil
}

// UpdateBook обновляет существующую книгу
func (r *BookRepository) UpdateBook(id int, book domain.Book) error {
	_, err := r.DB.Exec("UPDATE books SET title=$1, description=$2, author=$3, published=$4, price=$5 WHERE id=$6",
		book.Title, book.Description, book.Author, book.Published, book.Price, id)
	if err != nil {
		r.Log.WithError(err).Errorf("Failed to update book with ID: %d", id)
	}

	return err
}

// DeleteBook удаляет книгу по ID
func (r *BookRepository) DeleteBook(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		r.Log.WithError(err).Errorf("Failed to delete book with ID: %d", id)
	}
	return err
}
