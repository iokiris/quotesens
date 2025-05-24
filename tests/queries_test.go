package tests

import (
	"bscaut-test/internal/repository"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestInsertQuote(t *testing.T) {
	ctx := context.Background()
	godotenv.Load("../.env")
	pool, err := pgxpool.New(ctx, os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatalf("failed to connect to DB: %v", err)
	}
	defer pool.Close()

	repo := repository.NewQuoteRepository(pool)

	err = repo.AddQuote(ctx, "Test Author", "Test Quote")
	if err != nil {
		t.Fatalf("InsertQuote returned error: %v", err)
	}

	queries, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}
	t.Log(queries)

	id := queries[0].ID
	author := queries[0].Author

	t.Log(repo.GetRandom(ctx))
	t.Log(repo.GetByAuthor(ctx, author))

	t.Log(repo.DeleteByID(ctx, id))
}
