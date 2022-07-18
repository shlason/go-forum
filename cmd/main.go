package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/configs"
	"github.com/shlason/go-forum/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteAuthRoutes(r)
	routes.RegisteUserRoutes(r)
	routes.RegisteThreadRoutes(r)
	routes.RegistePostRoutes(r)

	http.ListenAndServe(configs.ServerCfg.Addr, r)
}
