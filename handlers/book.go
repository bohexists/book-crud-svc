package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"

	"github.com/bohexists/book-crud-svc/db"
	"github.com/bohexists/book-crud-svc/models"
)

// c is a global cache
var c = cache.New(5*time.Minute, 10*time.Minute)

// GetBooks retrieves all books from the database
func GetBooks(w http.ResponseWriter, r *http.Request) {

	// If the data is in the cache, return it
	if x, found := c.Get("books"); found {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(x); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If the data is not in the cache, query the database
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

	// Add the data to the cache
	c.Set("books", bookList, cache.DefaultExpiration)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetBook retrieves a single book from the database
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// If the data is in the cache, return it
	if x, found := c.Get("book_" + strconv.Itoa(id)); found {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(x); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If the data is not in the cache, query the database
	var book models.Book
	err = db.DB.QueryRow("SELECT id, title, author, published FROM books WHERE id=$1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Published)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Add the data to the cache
	c.Set("book_"+strconv.Itoa(id), book, cache.DefaultExpiration)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
