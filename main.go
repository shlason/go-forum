package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/shlason/go-forum/docs"
	"github.com/shlason/go-forum/pkg/configs"
	"github.com/shlason/go-forum/pkg/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteAuthRoutes(r)
	routes.RegisteUserRoutes(r)
	routes.RegisteThreadRoutes(r)
	routes.RegistePostRoutes(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(configs.ServerCfg.Addr, r)
}
