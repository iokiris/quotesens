package server

import (
	"bscaut-test/internal/middleware"
	"net/http"
)

type Route struct {
	Pattern     string
	Handler     middleware.HandlerFuncWithError
	Middlewares []func(http.Handler) http.Handler
	Methods     []string
}

func NewRouter(mux *http.ServeMux, routes []Route) *http.ServeMux {
	for _, route := range routes {
		var h http.Handler = middleware.ErrorHandler(route.Handler) // error handler вместо обычного хттп

		for i := len(route.Middlewares) - 1; i >= 0; i-- {
			h = route.Middlewares[i](h)
		}

		handlerWithMethod := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowed := false
			for _, m := range route.Methods {
				if m == r.Method {
					allowed = true
				}
			}
			if !allowed {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			h.ServeHTTP(w, r)
		})
		mux.Handle(route.Pattern, handlerWithMethod)
	}
	return mux
}
