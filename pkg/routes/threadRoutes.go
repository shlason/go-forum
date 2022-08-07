package routes

import (
	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/controllers"
	"github.com/shlason/go-forum/pkg/middlewares"
)

func RegisteThreadRoutes(router *mux.Router) {
	router.Handle("/threads", middlewares.Adapt(controllers.Thread.GetThreads, middlewares.Header())).Methods("GET")
	router.Handle("/threads", middlewares.Adapt(controllers.Thread.CreateThread, middlewares.Auth(), middlewares.Header())).Methods("POST")
	router.Handle("/threads/{threadID}", middlewares.Adapt(controllers.Thread.GetThreadById, middlewares.Header())).Methods("GET")
	router.Handle("/threads/{threadID}", middlewares.Adapt(controllers.Thread.UpdateThread, middlewares.Auth(), middlewares.Header())).Methods("PATCH")
	router.Handle("/threads/{threadID}/posts", middlewares.Adapt(controllers.Thread.GetThreadRelatedPosts, middlewares.Header())).Methods("GET")
	router.Handle("/threads/{threadID}/posts", middlewares.Adapt(controllers.Thread.CreateThreadRelatedPost, middlewares.Header())).Methods("POST")
}
