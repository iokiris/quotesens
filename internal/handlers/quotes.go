package handlers

import (
	"bscaut-test/internal/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) AddQuote(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return nil
	}
	var input struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding body:", err)
		return err
	}
	err := h.service.AddQuote(r.Context(), input.Author, input.Quote)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *Handler) GetAllQuotes(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return nil
	}
	author := r.URL.Query().Get("author")
	var quotes []model.Quote
	var err error

	if author != "" {
		quotes, err = h.service.GetByAuthor(r.Context(), author)
	} else {
		quotes, err = h.service.GetAll(r.Context())
	}

	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(quotes)
	return nil
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return nil
	}
	quote, err := h.service.GetRandom(r.Context())
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(quote)
	return nil
}

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "DELETE" {
		http.Error(w, "Only DELETE method is supported.", http.StatusMethodNotAllowed)
		return nil
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = h.service.DeleteByID(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
