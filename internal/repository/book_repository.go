package repository

import (
	"database/sql"
	"github.com/bohexists/book-crud-svc/internal/domain"
	"log"
)

// BookRepository is a repository for books
type BookRepository struct {
	db *sql.DB
}

// NewBookRepository creates a new BookRepository
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

// GetBooks retrieves all books from the database
func (r *BookRepository) GetBooks() ([]domain.Book, error) {
	books := []domain.Book{}
	// Execute the query
	query := "SELECT id, title, description, author, published, price FROM books"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Close the rows when the function returns
	defer rows.Close()
	// Iterate through the rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Published, &book.Price)
		if err != nil {
			log.Println("Error scanning book:", err)
			continue
		}
		// Add the book to the list
		books = append(books, book)
	}
	// Check for any errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Return the list of books
	return books, nil
}
