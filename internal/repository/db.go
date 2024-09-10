package repository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DB is the global database connection
var db *sql.DB

func SetupDatabase() *sql.DB {

	// Connect to testing database if in testing mode
	envFile := ".env"
	if testingMode := os.Getenv("TESTING"); testingMode == "true" {
		envFile = ".env.test"
	}

	// Load environment variables
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// Open the connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// Print a success message
	fmt.Println("Database connection established")

	return db
}
