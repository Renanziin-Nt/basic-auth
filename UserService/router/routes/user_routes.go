package routes

import (
	"net/http"
)

var routesUser = []Routes{
	{
		Uri:              "/login",
		Method:           http.MethodPost,
		Controller:       func(w http.ResponseWriter, r *http.Request) {},
		ReqAuthorization: false,
	},
	{
		Uri:              "/register",
		Method:           http.MethodPost,
		Controller:       func(w http.ResponseWriter, r *http.Request) {},
		ReqAuthorization: false,
	},
}
