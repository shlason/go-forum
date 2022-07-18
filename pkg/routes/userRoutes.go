package routes

import "github.com/gorilla/mux"

func RegisteUserRoutes(router *mux.Router) {
	router.Handle("/users/info", tempHandler).Methods("GET")
	router.Handle("/users/info", tempHandler).Methods("PATCH")
}
