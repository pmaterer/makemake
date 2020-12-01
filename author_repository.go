package main

import "database/sql"

// AuthorRepository ...
type AuthorRepository struct {
	db *sql.DB
}

// NewAuthorRepository ...
func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db}
}

// AddAuthor ...
func (r *AuthorRepository) AddAuthor(a Author) (int64, error) {

	stmt, err := r.db.Prepare("INSERT INTO authors (firstname, lastname) VALUES ($1, $2)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(a.Firstname, a.Lastname)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil

}

// GetAllAuthors ...
func (r *AuthorRepository) GetAllAuthors() ([]Author, error) {
	var authors []Author

	rows, err := r.db.Query("SELECT id, firstname, lastname FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Author
		err := rows.Scan(&a.ID, &a.Firstname, &a.Lastname)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}

// GetAuthor ...
func (r *AuthorRepository) GetAuthor(id int64) (Author, error) {
	var author Author
	err := r.db.QueryRow("SELECT id, firstname, lastname FROM authors WHERE id = $1", id).Scan(&author.ID, &author.Firstname, &author.Lastname)
	if err != nil {
		return author, err
	}
	return author, nil
}
