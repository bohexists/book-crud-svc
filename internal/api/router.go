package api

import (
	"github.com/bohexists/book-crud-svc/internal/middleware"
	"github.com/bohexists/book-crud-svc/internal/service"
	"github.com/gorilla/mux"
)

func NewRouter(bookService *service.BookService) *mux.Router {
	// Create a new router
	var r *mux.Router
	r = mux.NewRouter()

	// Create a new BookHandler
	bookHandler := &BookHandler{Service: bookService}

	r.HandleFunc("/login", middleware.Login).Methods("POST")
	r.HandleFunc("/books", middleware.AuthenticateJWT(bookHandler.GetBooks)).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(bookHandler.GetBook)).Methods("GET")
	r.HandleFunc("/books", middleware.AuthenticateJWT(bookHandler.CreateBook)).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(bookHandler.UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(bookHandler.DeleteBook)).Methods("DELETE")

	return r
}
