package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/shlason/go-forum/docs"
	"github.com/shlason/go-forum/pkg/configs"
	"github.com/shlason/go-forum/pkg/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Go-forum API
// @version         1.0
// @description     This is a sample forum server.

// @contact.email  nocvi111@gmail.com

// @license.name  MIT
// @license.url   https://github.com/shlason/go-forum/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /
func main() {
	r := mux.NewRouter()
	routes.RegisteAuthRoutes(r)
	routes.RegisteUserRoutes(r)
	routes.RegisteThreadRoutes(r)
	routes.RegistePostRoutes(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(configs.ServerCfg.Addr, r)
}
