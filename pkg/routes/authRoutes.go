package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/controllers"
	"github.com/shlason/go-forum/pkg/middlewares"
)

// TODO: All handlers

type customHandler func(http.ResponseWriter, *http.Request)

func (ch customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch(w, r)
}

var tempHandler = customHandler(func(w http.ResponseWriter, r *http.Request) {})

func RegisteAuthRoutes(router *mux.Router) {
	router.Handle("/signup", middlewares.Adapt(controllers.Auth.Signup, middlewares.Header())).Methods("POST")
	router.Handle("/login", middlewares.Adapt(controllers.Auth.Login, middlewares.Header())).Methods("POST")
	router.Handle("/logout", middlewares.Adapt(controllers.Auth.Logout, middlewares.Auth(), middlewares.Header())).Methods("POST")
}
