package router

import (
	"net/http"

	"ticket-wallet/presentation/controller"

	"github.com/gorilla/mux"
)

const (
	apiVer = "/api/v1"
)

func NewRouter(cont *controller.Controller) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	{
		const routeName = "get-layout"
		router.Methods(http.MethodGet).
			Name(routeName).
			PathPrefix(apiVer).
			Handler(http.HandlerFunc(cont.GetHallLayout))
	}

	return router
}
