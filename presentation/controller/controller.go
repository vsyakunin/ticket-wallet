package controller

import (
	"log"
	"net/http"
)

type Controller struct {
	Svc Service
}

func NewController(svc Service) *Controller {
	return &Controller{Svc: svc}
}

func (c *Controller) GetHallLayout(w http.ResponseWriter, r *http.Request) {
	hallLayout, err := c.Svc.GetHallLayout()
	if err != nil {
		log.Println(getLogMessage(err, r.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, hallLayout) {
		log.Println("error while writing")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
