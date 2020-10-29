package repository

import (
	"database/sql"
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/pmaterer/makemake/domain"
	"github.com/stretchr/testify/assert"
)

var b = domain.Book{
	ID:     66,
	Title:  "The Stand",
	Author: "Stephen King",
	ISBN:   "123456789",
}

var b2 = domain.Book{
	ID:     67,
	Title:  "IT",
	Author: "Stephen King",
	ISBN:   "1011213",
}

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating the sql db mock: %s", err)
	}

	return db, mock
}

func TestAddBook(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	mock.ExpectExec("INSERT INTO books").
		WithArgs(b.Title, b.Author, b.ISBN).WillReturnResult(sqlmock.NewResult(65, 1))
	repo := CreateBookRepository(db)
	err := repo.AddBook(b)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestDeleteBook(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	mock.ExpectExec("DELETE FROM books").
		WithArgs(b.ID).WillReturnResult(sqlmock.NewResult(65, 1))
	repo := CreateBookRepository(db)
	err := repo.DeleteBook(b.ID)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetBook(t *testing.T) {
	db, mock := newMock()
	defer db.Close()
	repo := CreateBookRepository(db)

	query := "SELECT id, title, author, isbn FROM books WHERE id = \\$1"
	rows := sqlmock.NewRows([]string{"id", "title", "author", "isbn"}).
		AddRow(b.ID, b.Title, b.Author, b.ISBN)

	mock.ExpectQuery(query).WithArgs(b.ID).WillReturnRows(rows)

	book, err := repo.GetBook(b.ID)
	assert.NotNil(t, book)
	assert.NoError(t, err)
}

func TestGetAllBooks(t *testing.T) {
	db, mock := newMock()
	defer db.Close()
	repo := CreateBookRepository(db)

	query := "SELECT id, title, author, isbn FROM books"
	rows := sqlmock.NewRows([]string{"id", "title", "author", "isbn"}).
		AddRow(b.ID, b.Title, b.Author, b.ISBN).
		AddRow(b2.ID, b2.Title, b2.Author, b2.ISBN)

	mock.ExpectQuery(query).WillReturnRows(rows)

	books, err := repo.GetAllBooks()
	assert.NotNil(t, books)
	assert.NoError(t, err)
	assert.Equal(t, len(books), 2, "they should be equal")
}

func TestUpdateBook(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	mock.ExpectExec("UPDATE books SET title=\\$1, author=\\$2, isbn=\\$3 WHERE id=\\$4").
		WithArgs(b.Title, b.Author, b.ISBN, b.ID).WillReturnResult(sqlmock.NewResult(65, 1))
	repo := CreateBookRepository(db)
	err := repo.UpdateBook(b)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
