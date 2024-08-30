package main

import (
	_ "github.com/lib/pq"

	"github.com/yourusername/book-crud-svc/db"
)

func main() {

	db.InitDB()

}
