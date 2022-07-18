package routes

import (
	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/middlewares"
)

func RegistePostRoutes(router *mux.Router) {
	router.Handle("/posts", middlewares.Adapt(tempHandler, middlewares.Header())).Methods("GET")
	router.Handle("/posts/{postUUID}", middlewares.Adapt(tempHandler, middlewares.Header())).Methods("GET")
	router.Handle("/posts/{postUUID}", middlewares.Adapt(tempHandler, middlewares.Auth(), middlewares.Header())).Methods("PATCH")
}
