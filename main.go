package main

import (
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
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBook).Methods("DELETE")

	// Start the server
	log.Println("Server is running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))

}
