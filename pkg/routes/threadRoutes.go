package routes

import "github.com/gorilla/mux"

func RegisteThreadRoutes(router *mux.Router) {
	router.Handle("/threads", tempHandler).Methods("GET")
	router.Handle("/threads/{threadID}", tempHandler).Methods("GET")
	router.Handle("/threads/{threadID}", tempHandler).Methods("PATCH")
	router.Handle("/threads/{threadID}/posts", tempHandler).Methods("GET")
}
