package service

import (
	"bscaut-test/internal/exceptions"
	"bscaut-test/internal/model"
	"bscaut-test/internal/repository"
	"context"
	"strings"
)

type QuoteService struct {
	repo *repository.QuoteRepository
}

type QuoteServiceInterface interface {
	AddQuote(ctx context.Context, author, quote string) error
	GetAll(ctx context.Context) ([]model.Quote, error)
	GetByAuthor(ctx context.Context, author string) ([]model.Quote, error)
	GetRandom(ctx context.Context) (model.Quote, error)
	DeleteByID(ctx context.Context, id int) error
}

func NewService(repo *repository.QuoteRepository) *QuoteService {
	return &QuoteService{
		repo: repo,
	}
}

// AddQuote добавляет цитату с валидацией
func (s *QuoteService) AddQuote(ctx context.Context, author, quote string) error {
	author = strings.TrimSpace(author)
	quote = strings.TrimSpace(quote)

	if author == "" {
		return exceptions.Err(exceptions.ErrTypeInvalidInput, "Должен быть указан автор.", nil)
	}
	if quote == "" {
		return exceptions.Err(exceptions.ErrTypeInvalidInput, "Должна быть указана цитата.", nil)
	}
	return s.repo.AddQuote(ctx, author, quote)
}

// GetAll возвращает все цитаты
func (s *QuoteService) GetAll(ctx context.Context) ([]model.Quote, error) {
	return s.repo.GetAll(ctx)
}

// GetRandom возвращает случайную цитату
func (s *QuoteService) GetRandom(ctx context.Context) (model.Quote, error) {
	return s.repo.GetRandom(ctx)
}

// GetByAuthor возвращает цитаты по автору
func (s *QuoteService) GetByAuthor(ctx context.Context, author string) ([]model.Quote, error) {
	author = strings.TrimSpace(author)
	if author == "" {
		return nil, exceptions.Err(exceptions.ErrTypeInvalidInput, "Должен быть указан автор.", nil)
	}
	return s.repo.GetByAuthor(ctx, author)
}

// DeleteByID удаление по айди с проверкой
func (s *QuoteService) DeleteByID(ctx context.Context, id int) error {
	if id <= 0 {
		return exceptions.Err(exceptions.ErrTypeInvalidInput, "ID должен быть положительным числом", nil)
	}
	return s.repo.DeleteByID(ctx, id)
}
