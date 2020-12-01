package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pmaterer/makemake/interfaces/handlers"
	"github.com/pmaterer/makemake/mocks"
)

var bookService *mocks.Service

func init() {
	bookService = mocks.NewMockBookService()
}

func TestAddBook(t *testing.T) {
	jsonBody := strings.NewReader(`{"id":0,"title":"The Stand","author":"Stephen King","isbn":"9780340951446"}`)
	req, err := http.NewRequest("POST", "/books", jsonBody)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h := handlers.CreateBookHandler(bookService)
	handler := http.HandlerFunc(h.AddBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("AddBook HTTP handler returned wrong status code: got %v, want %v",
			status, http.StatusCreated)
	}
}

func TestDeleteBook(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/books/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h := handlers.CreateBookHandler(bookService)
	handler := http.HandlerFunc(h.DeleteBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AddBook HTTP handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)
	}
}

func TestGetAllBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a ResponseRecorder which satisfies http.ResponseWriter and records the response
	rr := httptest.NewRecorder()

	h := handlers.CreateBookHandler(bookService)
	handler := http.HandlerFunc(h.GetAllBooks)

	// handler satisfies http.Handler so we can call the ServeHTTP method
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetBook HTTP handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)
	}
}

func TestGetBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/books/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a ResponseRecorder which satisfies http.ResponseWriter and records the response
	rr := httptest.NewRecorder()

	h := handlers.CreateBookHandler(bookService)
	handler := http.HandlerFunc(h.GetBook)

	// handler satisfies http.Handler so we can call the ServeHTTP method
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetBook HTTP handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)
	}

	expected := `{"id":0,"title":"The Stand","author":"Stephen King","isbn":"9780340951446"}`
	t.Log(expected)
	t.Log(rr.Body.String())
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("GetBook HTTP handler returned unexpected body: %v, want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateBook(t *testing.T) {
	jsonBody := strings.NewReader(`{"id":0,"title":"The Stand: Special Edition","author":"Stephen King","isbn":"9780340951442"}`)
	req, err := http.NewRequest("PUT", "/books", jsonBody)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h := handlers.CreateBookHandler(bookService)
	handler := http.HandlerFunc(h.UpdateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AddBook HTTP handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)
	}
}
