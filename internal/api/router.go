// Package main - Book Management API
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
	"github.com/bohexists/book-crud-svc/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/bohexists/book-crud-svc/docs"

	"github.com/gorilla/mux"
)

func NewRouter(bookService *service.BookService) *mux.Router {
	// Create a new router
	var r *mux.Router
	r = mux.NewRouter()

	// Create a new BookHandler
	bookHandler := &BookHandler{Service: bookService}

	// Serve swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Register routes
	r.HandleFunc("/login", middleware.Login).Methods("POST")
	r.HandleFunc("/books", middleware.AuthenticateJWT(bookHandler.GetBooks)).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(bookHandler.GetBook)).Methods("GET")
	r.HandleFunc("/books", middleware.AuthenticateJWT(bookHandler.CreateBook)).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(bookHandler.UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(bookHandler.DeleteBook)).Methods("DELETE")

	return r
}
