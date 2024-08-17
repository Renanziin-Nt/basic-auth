package router

import (
	"User_Service/router/routes"

	"github.com/gorilla/mux"
)

// Gerar retornará um router com as rotas configuradas
func RoutersGenerator() *mux.Router {
	r := mux.NewRouter()
	return routes.RoutesConfiguration(r)
}
