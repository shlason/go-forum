package routes

import "github.com/gorilla/mux"

func RegistePostRoutes(router *mux.Router) {
	router.Handle("/posts", tempHandler).Methods("GET")
	router.Handle("/posts/{postUUID}", tempHandler).Methods("GET")
	router.Handle("/posts/{postUUID}", tempHandler).Methods("PATCH")
}
