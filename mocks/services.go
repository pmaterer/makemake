package mocks

import (
	"errors"

	"github.com/pmaterer/makemake/domain"
)

type Service struct {
	r         *Repo
	ShouldErr bool
}

func NewMockBookService() *Service {
	service := &Service{
		r: NewMockBookRepo(),
	}
	return service
}

func (s *Service) AddBook(b domain.Book) error {
	if s.ShouldErr {
		return errors.New("")
	}
	return nil
}

func (s *Service) DeleteBook(id int) error {
	if s.ShouldErr {
		return errors.New("")
	}
	return nil
}

func (s *Service) GetAllBooks() ([]domain.Book, error) {
	if s.ShouldErr {
		return nil, errors.New("")
	}
	return s.r.Books, nil
}

func (s *Service) GetBook(id int) (domain.Book, error) {
	var b domain.Book
	if s.ShouldErr {
		return b, errors.New("")
	}
	return s.r.Books[0], nil
}

func (s *Service) UpdateBook(b domain.Book) error {
	if s.ShouldErr {
		return errors.New("")
	}
	return nil
}
