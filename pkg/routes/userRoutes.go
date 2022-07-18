package routes

import (
	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/middlewares"
)

func RegisteUserRoutes(router *mux.Router) {
	router.Handle("/users/info", middlewares.Adapt(tempHandler, middlewares.Header())).Methods("GET")
	router.Handle("/users/info", middlewares.Adapt(tempHandler, middlewares.Auth(), middlewares.Header())).Methods("PATCH")
}
