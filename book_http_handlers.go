package main

import (
	"encoding/json"
	"net/http"
)

// BookHTTPHandler ...
type BookHTTPHandler struct {
	r BookRepositoryHandler
}

// NewBookHTTPHandler ...
func NewBookHTTPHandler(r BookRepositoryHandler) BookHTTPHandler {
	return BookHTTPHandler{r}
}

func (h BookHTTPHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var a Author
	var b Book
	b.Author = a

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
	}
	respondWithJSON(w, http.StatusOK, b)
}
