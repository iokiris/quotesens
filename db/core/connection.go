package db

import (
	"bscaut-test/pkg/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func NewPgxPool(cfg *config.Config) *pgxpool.Pool {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresDB,
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal("(Panic): DB connect error: ", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("(Panic): DB ping error: ", err)
	}

	log.Println("Connected to DB. Pool Created.")

	return pool
}
