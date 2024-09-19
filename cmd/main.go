package main

import (
	"github.com/bohexists/book-crud-svc/internal/api"
	"github.com/bohexists/book-crud-svc/internal/repository"
	"github.com/bohexists/book-crud-svc/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"net/http"
	"os"

	_ "github.com/bohexists/book-crud-svc/docs"
	_ "github.com/lib/pq"
)

var log = logrus.New()

func init() {
	// Настройка logrus для вывода в JSON
	log.Formatter = &logrus.JSONFormatter{}

	// Уровень логирования можно настроить через переменную окружения
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Level = logrus.InfoLevel // Уровень по умолчанию
	} else {
		log.Level = logLevel
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatal("Failed to load .env file")
	}

	// Initialize the logger
	log = logrus.New()
	// Initialize the database connection
	db := repository.SetupDatabase()
	// Create a new repository for books
	bookRepo := repository.NewBookRepository(db, log)
	// Create a new repository for users (used for authentication)
	userRepo := repository.NewUserRepository(db)
	// Create a new service for books
	bookService := service.NewBookService(bookRepo)
	// Create a new router
	router := api.NewRouter(bookService, userRepo)

	// Start the server
	log.Println("Server is running on port", os.Getenv("PORT"))
	err = (http.ListenAndServe(":"+os.Getenv("PORT"), router))
	if err != nil {
		log.WithError(err).Fatal("Server failed to start")
	}

}
