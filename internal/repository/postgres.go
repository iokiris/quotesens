package repository

import (
	"bscaut-test/db/queries"
	"bscaut-test/internal/exceptions"
	"bscaut-test/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuoteRepository struct {
	pool *pgxpool.Pool
}

func NewQuoteRepository(pool *pgxpool.Pool) *QuoteRepository {
	return &QuoteRepository{
		pool: pool,
	}
}

// AddQuote добавление цитаты в бд
func (r *QuoteRepository) AddQuote(ctx context.Context, author, quote string) error {
	_, err := r.pool.Exec(ctx, queries.AddNewQuote, author, quote)
	return err
}

// GetAll возвращает все цитаты группой (id, author, quote)
func (r *QuoteRepository) GetAll(ctx context.Context) ([]model.Quote, error) {
	var quotes []model.Quote
	rows, err := r.pool.Query(ctx, queries.GetAllQuotesGrouped)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var q model.Quote
		if err := rows.Scan(&q.ID, &q.Author, &q.Quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return quotes, nil
}

// GetRandom случайная цитата
func (r *QuoteRepository) GetRandom(ctx context.Context) (model.Quote, error) {
	query := queries.GetRandomQuoteGrouped
	var q model.Quote
	err := r.pool.QueryRow(ctx, query).Scan(&q.ID, &q.Author, &q.Quote)
	if errors.Is(err, pgx.ErrNoRows) {
		return q, fmt.Errorf("no quotes found")
	}
	return q, err
}

// GetByAuthor фильтрация по автору
func (r *QuoteRepository) GetByAuthor(ctx context.Context, author string) ([]model.Quote, error) {
	query := queries.GetByAuthorGrouped
	rows, err := r.pool.Query(ctx, query, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []model.Quote
	for rows.Next() {
		var q model.Quote
		if err := rows.Scan(&q.ID, &q.Author, &q.Quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}
	return quotes, rows.Err()
}

// DeleteByID удаление цитаты по айди
func (r *QuoteRepository) DeleteByID(ctx context.Context, id int) error {
	query := queries.DeleteById
	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return exceptions.Err(exceptions.ErrTypeNotFound, "Цитата с таким ID не найдена.", err)
	}
	return nil
}
