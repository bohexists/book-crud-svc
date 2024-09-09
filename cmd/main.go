package main

import (
	"github.com/bohexists/book-crud-svc/auth"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/bohexists/book-crud-svc/db"
	"github.com/bohexists/book-crud-svc/handlers"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db.InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/login", auth.Login).Methods("POST")                                              // Login route
	r.HandleFunc("/books", auth.AuthenticateJWT(handlers.GetBooks)).Methods("GET")                  // Get books route
	r.HandleFunc("/books/{id:[0-9]+}", auth.AuthenticateJWT(handlers.GetBook)).Methods("GET")       // Get book route
	r.HandleFunc("/books", auth.AuthenticateJWT(handlers.CreateBook)).Methods("POST")               // Create book route
	r.HandleFunc("/books/{id:[0-9]+}", auth.AuthenticateJWT(handlers.UpdateBook)).Methods("PUT")    // Update book route
	r.HandleFunc("/books/{id:[0-9]+}", auth.AuthenticateJWT(handlers.DeleteBook)).Methods("DELETE") // Delete book route

	// Start the server
	log.Println("Server is running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))

}
