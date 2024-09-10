package repository

import (
	"database/sql"
	"github.com/bohexists/book-crud-svc/internal/domain"
)

// BookRepository структура для работы с базой данных по книгам
type BookRepository struct {
	DB *sql.DB
}

// NewBookRepository создает новый экземпляр BookRepository
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

// GetBooks извлекает все книги из базы данных
func (r *BookRepository) GetBooks() ([]domain.Book, error) {
	rows, err := r.DB.Query("SELECT id, title, description, author, published, price FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Published, &book.Price); err != nil {
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
		return domain.Book{}, err
	}
	return book, nil
}

// CreateBook добавляет новую книгу в базу данных
func (r *BookRepository) CreateBook(book domain.Book) (domain.Book, error) {
	err := r.DB.QueryRow("INSERT INTO books (title, description, author, published, price) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		book.Title, book.Description, book.Author, book.Published, book.Price).Scan(&book.ID)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

// UpdateBook обновляет существующую книгу
func (r *BookRepository) UpdateBook(id int, book domain.Book) error {
	_, err := r.DB.Exec("UPDATE books SET title=$1, description=$2, author=$3, published=$4, price=$5 WHERE id=$6",
		book.Title, book.Description, book.Author, book.Published, book.Price, id)
	return err
}

// DeleteBook удаляет книгу по ID
func (r *BookRepository) DeleteBook(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
