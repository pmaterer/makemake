package repository

import (
	"database/sql"

	"github.com/pmaterer/makemake/domain"
)

func CreateBookRepository(db *sql.DB) domain.BookRepository {
	return repository{db}
}

type repository struct {
	db *sql.DB
}

func (r repository) AddBook(b domain.Book) error {
	_, err := r.db.Exec(`INSERT INTO books(title, author, isbn) VALUES($1, $2, $3)`, b.Title, b.Author, b.ISBN)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) DeleteBook(id int) error {
	_, err := r.db.Exec(`DELETE FROM books WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) GetAllBooks() ([]domain.Book, error) {

	books := []domain.Book{}

	rows, err := r.db.Query(`SELECT id, title, author, isbn FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b domain.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN)
		if err != nil {
			return books, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r repository) GetBook(id int) (domain.Book, error) {
	var b domain.Book
	err := r.db.QueryRow(`SELECT id, title, author, isbn FROM books WHERE id = $1`, id).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r repository) UpdateBook(b domain.Book) error {
	_, err := r.db.Exec(`UPDATE books SET title=$1, author=$2, isbn=$3 WHERE id=$4`, b.Title, b.Author, b.ISBN, b.ID)
	if err != nil {
		return err
	}
	return nil
}
