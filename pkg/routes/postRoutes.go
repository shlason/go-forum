package routes

import (
	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/controllers"
	"github.com/shlason/go-forum/pkg/middlewares"
)

func RegistePostRoutes(router *mux.Router) {
	router.Handle("/posts", middlewares.Adapt(controllers.Post.GetPosts, middlewares.Header())).Methods("GET")
	router.Handle("/posts/{postUUID}", middlewares.Adapt(controllers.Post.GetPostByID, middlewares.Header())).Methods("GET")
	router.Handle("/posts/{postUUID}", middlewares.Adapt(controllers.Post.UpdatePost, middlewares.Auth(), middlewares.Header())).Methods("PATCH")
}
