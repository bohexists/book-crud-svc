package main

import (
	"github.com/bohexists/book-crud-svc/internal/api"
	"github.com/bohexists/book-crud-svc/internal/repository"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	repository.InitDB()

	// Create a new router
	router := api.NewRouter()

	// Start the server
	log.Println("Server is running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
