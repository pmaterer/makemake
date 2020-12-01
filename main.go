package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "db"
	dbPort     = 5432
	dbUser     = "makemake"
	dbPassword = "changeit"
	dbName     = "makemake"
	dbSSLMode  = "disable"
)

func main() {

	db, err := connectPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// bookRepository := repository.CreateBookRepository(db)
	// bookService := services.CreateBookService(bookRepository)
	// bookHandler := handlers.CreateBookHandler(bookService)

	// authorRepo := NewAuthorRepository(db)
	bookRepo := NewBookRepository(db)

	bookHandler := NewBookHTTPHandler(bookRepo)
	r := mux.NewRouter()
	r.HandleFunc("/books", bookHandler.AddBook).Methods("POST")
	// r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
	// r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	// r.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	// r.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")

	log.Printf("Running server on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}

func connectPostgres() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
