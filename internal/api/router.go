package api

import (
	"github.com/bohexists/book-crud-svc/internal/middleware"
	"github.com/bohexists/book-crud-svc/internal/service"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	// Create a new router
	var r *mux.Router
	r = mux.NewRouter()

	r.HandleFunc("/login", middleware.Login).Methods("POST")                                             // Login route
	r.HandleFunc("/books", middleware.AuthenticateJWT(service.GetBooks)).Methods("GET")                  // Get books route
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(service.GetBook)).Methods("GET")       // Get book route
	r.HandleFunc("/books", middleware.AuthenticateJWT(service.CreateBook)).Methods("POST")               // Create book route
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(service.UpdateBook)).Methods("PUT")    // Update book route
	r.HandleFunc("/books/{id:[0-9]+}", middleware.AuthenticateJWT(service.DeleteBook)).Methods("DELETE") // Delete book route

	return r
}
