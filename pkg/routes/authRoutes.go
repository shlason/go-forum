package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: handlers
// TODO: auth, header middlewares

type customHandler func(http.ResponseWriter, *http.Request)

func (ch customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch(w, r)
}

var tempHandler = customHandler(func(w http.ResponseWriter, r *http.Request) {})

func RegisteAuthRoutes(router *mux.Router) {
	router.Handle("/signup", tempHandler).Methods("POST")
	router.Handle("/login", tempHandler).Methods("POST")
	router.Handle("/logout", tempHandler).Methods("POST")
}
