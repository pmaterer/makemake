package mocks

import (
	"errors"

	"github.com/pmaterer/makemake/domain"
)

type Repo struct {
	Books     []domain.Book
	ShouldErr bool
}

func (r *Repo) AddBook(b domain.Book) error {
	if r.ShouldErr {
		return errors.New("")
	}
	return nil
}

func (r *Repo) DeleteBook(id int) error {
	if r.ShouldErr {
		return errors.New("")
	}
	return nil
}

func (r *Repo) GetAllBooks() ([]domain.Book, error) {
	if r.ShouldErr {
		return nil, errors.New("")
	}
	return r.Books, nil
}

func (r *Repo) GetBook(id int) (domain.Book, error) {
	var b domain.Book
	if r.ShouldErr {
		return b, errors.New("")
	}
	return r.Books[id], nil
}

func (r *Repo) UpdateBook(b domain.Book) error {
	if r.ShouldErr {
		return errors.New("")
	}
	return nil
}

// NewMockBookRepo returns a Repo mock with seeded, mocked book data
func NewMockBookRepo() *Repo {
	repo := &Repo{
		Books: []domain.Book{
			{
				ID:     0,
				Title:  "The Stand",
				Author: "Stephen King",
				ISBN:   "9780340951446",
			},
			{
				ID:     1,
				Title:  "The Wind-Up Bird Chronicle",
				Author: "Haruki Murakami",
				ISBN:   "9780099448792",
			},
		},
	}
	return repo
}
