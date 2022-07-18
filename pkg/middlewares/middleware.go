package middlewares

import "net/http"

type middlewareHandler func(http.Handler) http.Handler

func Adapt(h http.Handler, middlewareHandlers ...middlewareHandler) http.Handler {
	for _, middlewareHandler := range middlewareHandlers {
		h = middlewareHandler(h)
	}
	return h
}
