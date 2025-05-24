package main

import (
	db "bscaut-test/db/core"
	"bscaut-test/internal/handlers"
	"bscaut-test/internal/middleware"
	"bscaut-test/internal/repository"
	"bscaut-test/internal/server"
	"bscaut-test/internal/service"
	"bscaut-test/pkg/config"
)

func run() error {
	cfg := config.Load()
	mainPool := db.NewPgxPool(cfg)
	defer mainPool.Close()

	repo := repository.NewQuoteRepository(mainPool)
	handler := handlers.NewHandler(service.NewService(repo))

	router := server.InitRouter(handler)
	noCorsRouter := middleware.CorsMiddleware(router) // чтобы работал index.html
	srv := server.New(cfg.ServerPort, noCorsRouter)

	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
