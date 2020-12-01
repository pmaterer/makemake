package services_test

import (
	"testing"

	"github.com/pmaterer/makemake/mocks"
	"github.com/pmaterer/makemake/services"
)

var bookRepo *mocks.Repo
var service services.BookService

func init() {
	bookRepo = mocks.NewMockBookRepo()
	service = services.CreateBookService(bookRepo)
}

func TestAddBook(t *testing.T) {

	bookRepo.ShouldErr = false
	err := service.AddBook(bookRepo.Books[0])
	if err != nil {
		t.Fatal(err)
	}

	bookRepo.ShouldErr = true
	err = service.AddBook(bookRepo.Books[0])
	if err == nil {
		t.Fatal(err)
	}
}

func TestDeleteBook(t *testing.T) {

	bookRepo.ShouldErr = false
	err := service.DeleteBook(0)
	if err != nil {
		t.Fatal(err)
	}
	bookRepo.ShouldErr = true
	err = service.DeleteBook(0)
	if err == nil {
		t.Fatal(err)
	}

}

func TestGetAllBooks(t *testing.T) {

	bookRepo.ShouldErr = false
	b, err := service.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != 2 {
		t.Fatalf("Expected returned list to be of size 2, got: %d", len(b))
	}

	bookRepo.ShouldErr = true
	_, err = service.GetAllBooks()
	if err == nil {
		t.Fatal(err)
	}
}

func TestGetBook(t *testing.T) {

	bookRepo.ShouldErr = false
	b, err := service.GetBook(0)
	if err != nil {
		t.Fatal(err)
	}
	if b.Author != "Stephen King" {
		t.Fatalf("Unexpected book author returned. Got: %s, wanted: %s", b.Title, "Stephen King")
	}

	bookRepo.ShouldErr = true
	_, err = service.GetBook(0)
	if err == nil {
		t.Fatal(err)
	}
}

func TestUpdateBook(t *testing.T) {

	bookRepo.ShouldErr = false
	err := service.UpdateBook(bookRepo.Books[0])
	if err != nil {
		t.Fatal(err)
	}

	bookRepo.ShouldErr = true
	err = service.UpdateBook(bookRepo.Books[0])
	if err == nil {
		t.Fatal(err)
	}
}
