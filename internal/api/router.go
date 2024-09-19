// Package api - Book Management API
// @title Book Management API
// @description This is a sample server for managing books.
// @version 1.0
// @BasePath /api/v1
// @schemes http https
// @host localhost:8080
// @ContactName Your Name
// @ContactEmail your_email@example.com

package api

import (
	"github.com/bohexists/book-crud-svc/internal/middleware"
	"github.com/bohexists/book-crud-svc/internal/repository"
	"github.com/bohexists/book-crud-svc/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/bohexists/book-crud-svc/docs"

	"github.com/gorilla/mux"
)

func NewRouter(bookService *service.BookService, userRepo *repository.UserRepository) *mux.Router {
	// Create a new router
	var r *mux.Router
	r = mux.NewRouter()

	// Create a new BookHandler
	bookHandler := &BookHandler{Service: bookService}

	authHandler := NewAuthHandler(userRepo)

	// Serve swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Register routes
	// Register public routes (no JWT required)
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Register protected routes (JWT required)
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	protected.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods("GET")
	protected.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	protected.HandleFunc("/books/{id:[0-9]+}", bookHandler.UpdateBook).Methods("PUT")
	protected.HandleFunc("/books/{id:[0-9]+}", bookHandler.DeleteBook).Methods("DELETE")

	return r
}
