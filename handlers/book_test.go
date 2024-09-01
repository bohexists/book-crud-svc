package handlers

import (
	"testing"

	"github.com/bohexists/book-crud-svc/db"
)

// TestMain sets up the database before running tests
func TestMain(m *testing.M) {
	// Initialize the database connection
	db.InitDB()

	// Start the tests
	m.Run()
}

// TestGetBooks tests the GetBooks function
//func TestGetBooks(t *testing.T) {
//	//Prepare the database
//	db.DB.Exec("DELETE FROM books")
//	db.DB.Exec("INSERT INTO books (title, author, published, description, price) VALUES ('Test Book 1', 'Author 1', '2023-01-01', 'Description 1', 100)")
//	db.DB.Exec("INSERT INTO books (title, author, published, description, price) VALUES ('Test Book 2', 'Author 2', '2023-02-01', 'Description 2', 200)")
//
//	// Create a new request
//	req, err := http.NewRequest("GET", "/books", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Set the request header
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(GetBooks)
//
//	// Call the handler
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//
//	// Check the response
//	expected := `{"books":[{"id":1,"title":"Test Book 1","description":"Description 1","author":"Author 1","published":"2023-01-01T00:00:00Z","price":100},{"id":2,"title":"Test Book 2","description":"Description 2","author":"Author 2","published":"2023-02-01T00:00:00Z","price":200}]}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
//	}
//}

// TestGetBook tests the GetBook function
//func TestGetBook(t *testing.T) {
//	// Prepare the database
//	db.DB.Exec("DELETE FROM books")
//	db.DB.Exec("INSERT INTO books (id, title, author, published, description, price) VALUES (1, 'Test Book', 'Test Author', '2023-01-01', 'Test Description', 100)")
//
//	// Create a new request
//	req, err := http.NewRequest("GET", "/books/1", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Create a response recorder
//	rr := httptest.NewRecorder()
//	r := mux.NewRouter()
//	r.HandleFunc("/books/{id:[0-9]+}", GetBook)
//	r.ServeHTTP(rr, req)
//
//	// Check the status code
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//
//	// Check the response
//	expected := `{"id":1,"title":"Test Book","description":"Test Description","author":"Test Author","published":"2023-01-01T00:00:00Z","price":100}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
//	}
//}

// TestCreateBook tests the CreateBook function
//func TestCreateBook(t *testing.T) {
//	// Prepare the database
//	db.DB.Exec("DELETE FROM books")
//
//	// Create a new request
//	body := `{"title":"New Book","author":"New Author","published":"2023-03-01T00:00:00Z","description":"New Description","price":150}`
//	req, err := http.NewRequest("POST", "/books", strings.NewReader(body))
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Create a response recorder
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(CreateBook)
//
//	// Call the handler
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//
//	// Check the response
//	expected := `{"id":1,"title":"New Book","description":"New Description","author":"New Author","published":"2023-03-01T00:00:00Z","price":150}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
//	}
//}

// TestUpdateBook tests the UpdateBook function
//func TestUpdateBook(t *testing.T) {
//	// Prepare the database
//	db.DB.Exec("DELETE FROM books")
//	db.DB.Exec("INSERT INTO books (id, title, author, published, description, price) VALUES (1, 'Old Book', 'Old Author', '2023-01-01', 'Old Description', 100)")
//
//	// Create a new request
//	body := `{"title":"Updated Book","author":"Updated Author","published":"2023-04-01T00:00:00Z","description":"Updated Description","price":200}`
//	req, err := http.NewRequest("PUT", "/books/1", strings.NewReader(body))
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	//	Create a response recorder
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(UpdateBook)
//
//	// Call the handler
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code
//	if status := rr.Code; status != http.StatusNoContent {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
//	}
//
//	// Check the response
//	var updatedBook models.Book
//	err = db.DB.QueryRow("SELECT id, title, author, published, description, price FROM books WHERE id = 1").Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Published, &updatedBook.Description, &updatedBook.Price)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	expectedTitle := "Updated Book"
//	if updatedBook.Title != expectedTitle {
//		t.Errorf("book title not updated: got %v want %v", updatedBook.Title, expectedTitle)
//	}
//}

// TestDeleteBook tests the DeleteBook function
//func TestDeleteBook(t *testing.T) {
//	// Prepare the database
//	db.DB.Exec("DELETE FROM books") // Очистка таблицы перед тестом
//	db.DB.Exec("INSERT INTO books (id, title, author, published, description, price) VALUES (1, 'Book to Delete', 'Author', '2023-01-01', 'Description', 100)")
//
//	// Create a new request
//	req, err := http.NewRequest("DELETE", "/books/1", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Create a response recorder
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(DeleteBook)
//
//	// Call the handler
//	handler.ServeHTTP(rr, req)
//
//	// Check the status code
//	if status := rr.Code; status != http.StatusNoContent {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
//	}
//
//	// Check the response
//	var deletedBook models.Book
//	err = db.DB.QueryRow("SELECT id, title FROM books WHERE id = 1").Scan(&deletedBook.ID, &deletedBook.Title)
//	if err == nil {
//		t.Errorf("book not deleted: found book with ID 1")
//	}
//}
