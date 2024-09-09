package domain

import "time"

// Book represents a book in the library
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Published   time.Time `json:"published"`
	Price       int       `json:"price"`
}

// BookList represents a list of books
type BookList struct {
	Books []Book `json:"books"`
}

// AddBook adds a book to the list
func (bl *BookList) AddBook(book Book) {
	bl.Books = append(bl.Books, book)
}
