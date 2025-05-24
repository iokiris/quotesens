package units

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bscaut-test/internal/handlers"
	"bscaut-test/internal/model"
)

type MockService struct{}

func (m *MockService) AddQuote(ctx context.Context, author, quote string) error {
	if author == "" || quote == "" {
		return errors.New("пустой автор или цитата")
	}
	return nil
}

func (m *MockService) GetAll(ctx context.Context) ([]model.Quote, error) {
	return []model.Quote{
		{ID: 1, Author: "Лев Толстой", Quote: "Все счастливые семьи похожи друг на друга."},
		{ID: 2, Author: "Фёдор Достоевский", Quote: "Быть слишком разумным — та же ошибка, что и быть слишком глупым."},
		{ID: 3, Author: "Тестер", Quote: "Тестовая цитата"},
	}, nil
}

func (m *MockService) GetByAuthor(ctx context.Context, author string) ([]model.Quote, error) {
	if author == "Толстой" {
		return []model.Quote{
			{ID: 1, Author: "Лев Толстой", Quote: "Все счастливые семьи похожи друг на друга."},
		}, nil
	}
	if author == "Тестер" {
		return []model.Quote{
			{ID: 3, Author: "Тестер", Quote: "Тестовая цитата"},
		}, nil
	}
	return nil, nil
}

func (m *MockService) GetRandom(ctx context.Context) (model.Quote, error) {
	return model.Quote{ID: 3, Author: "Тестер", Quote: "Тестовая цитата"}, nil
}

func (m *MockService) DeleteByID(ctx context.Context, id int) error {
	if id == 999 {
		return errors.New("цитата не найдена")
	}
	return nil
}

func setupHandler() *handlers.Handler {
	return handlers.NewHandler(&MockService{})
}

func TestAddQuote(t *testing.T) {
	t.Log("TestAddQuote: старт теста")

	h := setupHandler()
	payload := `{"author":"Лев Толстой","quote":"Начало — половина дела."}`
	t.Logf("Payload: %s", payload)

	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddQuote(w, req)

	t.Logf("Response status: %d", w.Code)
	if w.Code != http.StatusCreated {
		t.Fatalf("ожидали статус 201 Created, получили %d, тело ответа: %s", w.Code, w.Body.String())
	} else {
		t.Log("AddQuote выполнен успешно")
	}
}

func TestGetAllQuotes(t *testing.T) {
	t.Log("TestGetAllQuotes: старт теста")

	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	w := httptest.NewRecorder()

	h.GetAllQuotes(w, req)

	t.Logf("Response status: %d", w.Code)
	if w.Code != http.StatusOK {
		t.Fatalf("ожидали статус 200 OK, получили %d, тело ответа: %s", w.Code, w.Body.String())
	}

	var quotes []model.Quote
	err := json.NewDecoder(w.Body).Decode(&quotes)
	if err != nil {
		t.Fatalf("не удалось декодировать ответ: %v", err)
	}

	t.Logf("Получено %d цитат", len(quotes))
	if len(quotes) == 0 {
		t.Fatal("ожидали получить цитаты, получили 0")
	}

	for i, q := range quotes {
		t.Logf("Цитата #%d: Автор: %s, Текст: %s", i+1, q.Author, q.Quote)
	}
}

func TestGetByAuthor(t *testing.T) {
	t.Log("TestGetByAuthor: старт теста")

	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/quotes?author=Толстой", nil)
	w := httptest.NewRecorder()

	h.GetAllQuotes(w, req)

	t.Logf("Response status: %d", w.Code)
	if w.Code != http.StatusOK {
		t.Fatalf("ожидали статус 200 OK, получили %d, тело ответа: %s", w.Code, w.Body.String())
	}

	var quotes []model.Quote
	err := json.NewDecoder(w.Body).Decode(&quotes)
	if err != nil {
		t.Fatalf("не удалось декодировать ответ: %v", err)
	}

	if len(quotes) == 0 {
		t.Fatal("ожидали получить цитаты автора Толстой, но получили 0")
	}

	for _, q := range quotes {
		t.Logf("Проверяем автора цитаты: %s", q.Author)
		if q.Author != "Лев Толстой" {
			t.Fatalf("ожидали автора Лев Толстой, получили %s", q.Author)
		}
	}
	t.Log("TestGetByAuthor успешно завершён")
}

func TestGetRandomQuote(t *testing.T) {
	t.Log("TestGetRandomQuote: старт теста")

	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/quotes/random", nil)
	w := httptest.NewRecorder()

	h.GetRandomQuote(w, req)

	t.Logf("Response status: %d", w.Code)
	if w.Code != http.StatusOK {
		t.Fatalf("ожидали статус 200 OK, получили %d, тело ответа: %s", w.Code, w.Body.String())
	}

	var quote model.Quote
	err := json.NewDecoder(w.Body).Decode(&quote)
	if err != nil {
		t.Fatalf("не удалось декодировать ответ: %v", err)
	}

	t.Logf("Получена цитата: Автор: %s, Текст: %s", quote.Author, quote.Quote)
	if quote.ID == 0 {
		t.Fatal("ожидали ненулевой ID цитаты")
	}
	t.Log("TestGetRandomQuote успешно завершён")
}

func TestDeleteQuote(t *testing.T) {
	t.Log("TestDeleteQuote: старт теста")

	h := setupHandler()
	req := httptest.NewRequest(http.MethodDelete, "/quotes/1", nil)
	w := httptest.NewRecorder()

	h.DeleteQuote(w, req)

	t.Logf("Response status: %d", w.Code)
	if w.Code != http.StatusNoContent {
		t.Fatalf("ожидали статус 204 No Content, получили %d, тело ответа: %s", w.Code, w.Body.String())
	}
	t.Log("TestDeleteQuote успешно завершён")
}

func TestDeleteQuoteNotFound(t *testing.T) {
	t.Log("TestDeleteQuoteNotFound: старт теста")

	h := setupHandler()
	req := httptest.NewRequest(http.MethodDelete, "/quotes/999", nil)
	w := httptest.NewRecorder()

	h.DeleteQuote(w, req)

	t.Logf("Response status: %d", w.Code)
	if w.Code == http.StatusNoContent {
		t.Fatal("ожидали ошибочный статус, но получили 204 No Content")
	}
	t.Log("TestDeleteQuoteNotFound успешно завершён с ошибкой, как и ожидалось")
}
