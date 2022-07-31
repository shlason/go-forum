package routes

import (
	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/controllers"
	"github.com/shlason/go-forum/pkg/middlewares"
)

func RegisteUserRoutes(router *mux.Router) {
	router.Handle("/users/info", middlewares.Adapt(controllers.User.GetUsers, middlewares.Header())).Methods("GET")
	router.Handle("/users/info", middlewares.Adapt(controllers.User.PatchUser, middlewares.Auth(), middlewares.Header())).Methods("PATCH")
}
