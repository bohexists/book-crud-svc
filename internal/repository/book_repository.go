package repository

import (
	"database/sql"
	"github.com/bohexists/book-crud-svc/internal/domain"
	"log"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetBooks() ([]domain.Book, error) {
	books := []domain.Book{}

	query := "SELECT id, title, description, author, published, price FROM books"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Published, &book.Price)
		if err != nil {
			log.Println("Error scanning book:", err)
			continue
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
