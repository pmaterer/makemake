package main

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// BookRepositoryHandler ...
type BookRepositoryHandler interface {
	AddBook(b Book) (int64, error)
	DeleteBook(id int64) (int64, error)
	GetAllBooks() ([]Book, error)
	GetBook(id int64) (Book, error)
	UpdateBook(id int64, b Book) (int64, error)
}

type BookService interface {
	AddBook(b Book) (int64, error)
}
