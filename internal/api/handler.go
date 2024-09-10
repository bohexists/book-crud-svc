package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bohexists/book-crud-svc/internal/domain"
	"github.com/bohexists/book-crud-svc/internal/service"
	"github.com/gorilla/mux"
)

// BookHandler is a handler for books
type BookHandler struct {
	Service *service.BookService
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		Service: service,
	}
}

// GetBooks processes a request to get all books
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseWithJSON(w, books, http.StatusOK)
}

// GetBook processes a request to get a specific book
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := h.Service.GetBook(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	responseWithJSON(w, book, http.StatusOK)
}

// CreateBook processes a request to create a new book
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdBook, err := h.Service.CreateBook(book)
	if err != nil {
		http.Error(w, "Error creating book", http.StatusInternalServerError)
		return
	}
	responseWithJSON(w, createdBook, http.StatusCreated)
}

// UpdateBook processes a request to update a book
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = h.Service.UpdateBook(id, book)
	if err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteBook processes a request to delete a book
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteBook(id); err != nil {
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// responseWithJSON sends a JSON response with the given status code
func responseWithJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
