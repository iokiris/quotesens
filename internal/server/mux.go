package server

import (
	"bscaut-test/internal/handlers"
	"bscaut-test/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(handler *handlers.Handler) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/quotes", middleware.ErrorHandler(handler.GetAllQuotes)).Methods("GET")
	router.HandleFunc("/quotes", middleware.ErrorHandler(handler.AddQuote)).Methods("POST")
	router.HandleFunc("/quotes/random", middleware.ErrorHandler(handler.GetRandomQuote)).Methods("GET")
	router.HandleFunc("/quotes/{id}", middleware.ErrorHandler(handler.DeleteQuote)).Methods("DELETE")

	timeOutRouter := middleware.TimeoutMiddleware(router)
	return timeOutRouter
}
