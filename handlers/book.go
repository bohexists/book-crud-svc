package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/bohexists/book-crud-svc/db"
	"github.com/bohexists/book-crud-svc/models"
)

// GetBooks retrieves all books from the database
func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, author, published FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookList models.BookList

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Published); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bookList.AddBook(book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

// GetBook retrieves a single book from the database
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = db.DB.QueryRow("SELECT id, title, author, published FROM books WHERE id=$1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Published)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook creates a new book in the database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow("INSERT INTO books (title, author, published) VALUES ($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Published).Scan(&book.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates an existing book in the database
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("UPDATE books SET title=$1, author=$2, published=$3 WHERE id=$4",
		book.Title, book.Author, book.Published, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteBook deletes a book from the database
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
