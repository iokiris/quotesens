package handlers

import (
	"bscaut-test/internal/service"
)

type Handler struct {
	service service.QuoteServiceInterface
}

func NewHandler(s service.QuoteServiceInterface) *Handler {
	return &Handler{
		service: s,
	}
}
