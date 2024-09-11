# Book CRUD Service

## Description
Book CRUD Service is a REST API for managing books, implemented in Go. The service allows you to create, read, update, and delete book information in the database.

## Technologies
- **Go**: The primary programming language.
- **PostgreSQL**: Database management system.
- **Docker**: Used for containerizing the application and its dependencies.
- **Swagger**: Used for API documentation.

## Installation and Launch

### Requirements
- Go (version 1.22 or higher)
- Docker and Docker Compose

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone https://yourrepository.com/book-crud-svc.git
   cd book-crud-svc
   ```

2. **Launch using Docker Compose**
   ```bash
   docker-compose up --build
   ```

   This command will build and launch all necessary containers (including PostgreSQL and the application itself).

3. **Local launch**
   For a local launch, ensure your database is accessible and run:
   ```bash
   go run cmd/main.go
   ```

## Usage

After launching, the server is available at `http://localhost:8080`. Swagger documentation is available at `http://localhost:8080/swagger/index.html`.

## API Endpoints

- `POST /login` - User authentication.
- `GET /books` - Retrieve a list of all books.
- `GET /books/{id}` - Retrieve a book by ID.
- `POST /books` - Add a new book.
- `PUT /books/{id}` - Update book information.
- `DELETE /books/{id}` - Delete a book.
