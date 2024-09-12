package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/bohexists/book-crud-svc/internal/domain"
	"github.com/bohexists/book-crud-svc/internal/service"
	"github.com/gorilla/mux"
)

// BookHandler is a handler for books
type BookHandler struct {
	Service BookServiceInterface
	Log     *logrus.Logger
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(service *service.BookService, log *logrus.Logger) *BookHandler {
	return &BookHandler{
		Service: service,
		Log:     log,
	}
}

type BookServiceInterface interface {
	GetBooks() ([]domain.Book, error)
	GetBook(id int) (domain.Book, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id int, book domain.Book) error
	DeleteBook(id int) error
}

// GetBooks godoc
// @Summary Get all books
// @Description Get all books from the database
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} domain.Book
// @Router /books [get]
// @Security ApiKeyAuth
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetBooks()
	if err != nil {
		h.Log.WithError(err).Error("Failed to fetch books")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.responseWithJSON(w, books, http.StatusOK)
}

// GetBook godoc
// @Summary Get a book by ID
// @Description Get a specific book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} domain.Book
// @Router /books/{id} [get]
// @Security ApiKeyAuth
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.Log.WithError(err).Error("Invalid book ID")
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := h.Service.GetBook(id)
	if err != nil {
		h.Log.WithError(err).WithField("book_id", id).Error("Book not found")
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	h.responseWithJSON(w, book, http.StatusOK)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided information
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.Book true "Create Book"
// @Success 201 {object} domain.Book
// @Router /books [post]
// @Security ApiKeyAuth
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.Log.WithError(err).Error("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdBook, err := h.Service.CreateBook(book)
	if err != nil {
		h.Log.WithError(err).Error("Failed to create book")
		http.Error(w, "Error creating book", http.StatusInternalServerError)
		return
	}
	h.responseWithJSON(w, createdBook, http.StatusCreated)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update a book with specified ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body domain.Book true "Update Book"
// @Success 204
// @Router /books/{id} [put]
// @Security ApiKeyAuth
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.Log.WithError(err).Error("Invalid book ID")
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.Log.WithError(err).Error("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = h.Service.UpdateBook(id, book)
	if err != nil {
		h.Log.WithError(err).WithField("book_id", id).Error("Failed to update book")
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 204
// @Router /books/{id} [delete]
// @Security ApiKeyAuth
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.Log.WithError(err).Error("Invalid book ID")
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteBook(id); err != nil {
		h.Log.WithError(err).WithField("book_id", id).Error("Failed to delete book")
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// responseWithJSON sends a JSON response with the given status code
func (h *BookHandler) responseWithJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Log.WithError(err).Error("Failed to encode response to JSON")
	}
}
