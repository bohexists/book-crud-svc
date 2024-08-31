package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Published   time.Time `json:"published"`
	Price       int       `json:"price"`
}

type BookList struct {
	Books []Book `json:"books"`
}

func (bl *BookList) AddBook(book Book) {
	bl.Books = append(bl.Books, book)
}
