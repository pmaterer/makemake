package main

import "database/sql"



// BookRepository ...
type BookRepository struct {
	db *sql.DB
}

// NewBookRepository ...
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db}
}

// AddBook ...
func (r *BookRepository) AddBook(b Book) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO book (title, author_id) VALUES ($1, $2)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(b.Title, b.Author.ID)
	if err != nil {
		return -1, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastID, nil
}

// DeleteBook ...
func (r *BookRepository) DeleteBook(id int64) (int64, error) {
	stmt, err := r.db.Prepare("DELETE FROM books WHERE id=$1")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return -1, err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowCount, nil
}

// GetAllBooks ...
func (r *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	rows, err := r.db.Query("SELECT id, title, author_id FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author.ID)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// GetBook ...
func (r *BookRepository) GetBook(id int64) (Book, error) {
	var book Book
	err := r.db.QueryRow("SELECT id, title, author_id FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author.ID)
	if err != nil {
		return book, err
	}
	return book, nil
}

// UpdateBook ...
func (r *BookRepository) UpdateBook(id int64, b Book) (int64, error) {
	stmt, err := r.db.Prepare("UPDATE books SET title=$1 WHERE id=$2")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(b.Title, id)
	if err != nil {
		return -1, err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowCount, nil
}
