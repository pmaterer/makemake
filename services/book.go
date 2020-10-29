package services

import (
	"github.com/pmaterer/makemake/domain"
)

type BookService interface {
	AddBook(b domain.Book) error
	DeleteBook(id int) error
	GetAllBooks() ([]domain.Book, error)
	GetBook(id int) (domain.Book, error)
	UpdateBook(b domain.Book) error
}

func CreateBookService(r domain.BookRepository) BookService {
	return bookService{r}
}

type bookService struct {
	r domain.BookRepository
}

func (s bookService) AddBook(b domain.Book) error {
	err := s.r.AddBook(b)
	if err != nil {
		return err
	}
	return nil
}

func (s bookService) DeleteBook(id int) error {
	err := s.r.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

func (s bookService) GetAllBooks() ([]domain.Book, error) {
	books, err := s.r.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s bookService) GetBook(id int) (domain.Book, error) {
	book, err := s.r.GetBook(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s bookService) UpdateBook(b domain.Book) error {
	err := s.r.UpdateBook(b)
	if err != nil {
		return err
	}
	return nil
}
