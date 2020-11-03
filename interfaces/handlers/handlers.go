package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pmaterer/makemake/domain"
	"github.com/pmaterer/makemake/services"
)

type BookHandler interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

func CreateBookHandler(s services.BookService) BookHandler {
	return bookHandler{s}
}

type bookHandler struct {
	s services.BookService
}

func (h bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var b domain.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := h.s.AddBook(b)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, "Created book")
}

func (h bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, _ := strconv.Atoi(idRaw)
	err := h.s.DeleteBook(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, "Deleted book")
}

func (h bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.s.GetAllBooks()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(books)
}

func (h bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, _ := strconv.Atoi(idRaw)
	book, err := h.s.GetBook(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}

func (h bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, _ := strconv.Atoi(idRaw)
	var b domain.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	b.ID = id
	err := h.s.UpdateBook(b)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, "Book updated")
}
