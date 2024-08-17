package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Uri              string
	Method           string
	Controller       func(http.ResponseWriter, *http.Request)
	ReqAuthorization bool
}

func RoutesConfiguration(router *mux.Router) *mux.Router {
	routes := routesUser

	for _, rota := range routes {
		router.HandleFunc(rota.Uri, rota.Controller).Methods(rota.Method)
	}

	return router
}
