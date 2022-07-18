package middlewares

import (
	"fmt"
	"net/http"
)

func Auth() middlewareHandler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Auth middleware start")
			h.ServeHTTP(w, r)
			fmt.Println("Auth middleware end")
		})
	}
}
