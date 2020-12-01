package main

type Author struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// AuthorRepositoryHandler ...
type AuthorRepositoryHandler interface {
	AddAuthor(a Author) (int64, error)
	GetAllAuthors() ([]Author, error)
	GetAuthor(id int64) (Author, error)
}
