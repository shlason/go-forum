package routes

import (
	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/middlewares"
)

func RegisteThreadRoutes(router *mux.Router) {
	router.Handle("/threads", middlewares.Adapt(tempHandler, middlewares.Header())).Methods("GET")
	router.Handle("/threads/{threadID}", middlewares.Adapt(tempHandler, middlewares.Header())).Methods("GET")
	router.Handle("/threads/{threadID}", middlewares.Adapt(tempHandler, middlewares.Auth(), middlewares.Header())).Methods("PATCH")
	router.Handle("/threads/{threadID}/posts", middlewares.Adapt(tempHandler, middlewares.Header())).Methods("GET")
}
