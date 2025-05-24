package middleware

import (
	"bscaut-test/internal/exceptions"
	"log"
	"net/http"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

func ErrorHandler(h HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}

		if t, msg, ok := exceptions.IsPublic(err); ok {
			httpCode := exceptions.MapErrorTypeToHTTPCode(t)
			if httpCode != 500 {
				http.Error(w, msg, httpCode)
				return
			}
		}

		log.Printf("internal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
