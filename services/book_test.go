package services

import (
	"testing"

	"github.com/pmaterer/makemake/domain"
	"github.com/pmaterer/makemake/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBook(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:     1,
		Title:  "The Stand",
		Author: "Stephen King",
		ISBN:   "12345678",
	}

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("AddBook", mock.Anything).Return(nil).Once()
		s := CreateBookService(mockBookRepo)
		err := s.AddBook(mockBook)

		assert.NoError(t, err)

		mockBookRepo.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:     1,
		Title:  "The Stand",
		Author: "Stephen King",
		ISBN:   "12345678",
	}

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("DeleteBook", mock.AnythingOfType("int")).Return(nil).Once()
		s := CreateBookService(mockBookRepo)
		err := s.DeleteBook(mockBook.ID)

		assert.NoError(t, err)

		mockBookRepo.AssertExpectations(t)
	})
}

func TestGetBook(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:     1,
		Title:  "The Stand",
		Author: "Stephen King",
		ISBN:   "12345678",
	}

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("GetBook", mock.AnythingOfType("int")).Return(mockBook, nil).Once()
		s := CreateBookService(mockBookRepo)
		b, err := s.GetBook(mockBook.ID)

		assert.NoError(t, err)
		assert.NotNil(t, b)

		mockBookRepo.AssertExpectations(t)
	})
}

func TestGetAllBooks(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:     1,
		Title:  "The Stand",
		Author: "Stephen King",
		ISBN:   "12345678",
	}
	mockBooks := make([]domain.Book, 0)
	mockBooks = append(mockBooks, mockBook)

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("GetAllBooks").Return(mockBooks, nil).Once()
		s := CreateBookService(mockBookRepo)
		b, err := s.GetAllBooks()

		assert.NoError(t, err)
		assert.NotNil(t, b)

		mockBookRepo.AssertExpectations(t)
	})
}

func TestUpdateBook(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := domain.Book{
		ID:     1,
		Title:  "The Stand",
		Author: "Stephen King",
		ISBN:   "12345678",
	}

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("UpdateBook", mock.Anything).Return(nil).Once()
		s := CreateBookService(mockBookRepo)
		err := s.UpdateBook(mockBook)

		assert.NoError(t, err)

		mockBookRepo.AssertExpectations(t)
	})
}
