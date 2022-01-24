package router

import (
	"net/http"

	"github.com/vsyakunin/ticket-wallet/presentation/controller"

	"github.com/gorilla/mux"
)

const (
	apiVer = "/api/v1"
)

func NewRouter(cont *controller.Controller) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	{
		const routeName = "layout"
		router.Methods(http.MethodGet).
			Name(routeName).
			PathPrefix(apiVer).
			Path("/layout").
			Handler(http.HandlerFunc(cont.GetHallLayout))
	}

	{
		const routeName = "start-seating"
		router.Methods(http.MethodPost).
			Name(routeName).
			PathPrefix(apiVer).
			Path("/seating/start").
			Handler(http.HandlerFunc(cont.StartSeating))
	}

	{
		const routeName = "seating-results"
		router.Methods(http.MethodGet).
			Name(routeName).
			PathPrefix(apiVer).
			Path("/seating/results").
			Handler(http.HandlerFunc(cont.GetSeatingResults))
	}

	return router
}
