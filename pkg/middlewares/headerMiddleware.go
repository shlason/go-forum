package middlewares

import (
	"fmt"
	"net/http"
)

func Header() middlewareHandler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Header middleware start")
			h.ServeHTTP(w, r)
			fmt.Println("Header middleware end")
		})
	}
}
