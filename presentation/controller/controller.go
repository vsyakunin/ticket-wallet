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

func (c *Controller) StartSeating(w http.ResponseWriter, r *http.Request) {
	startSeatingPayload, err := extractStartSeatingPayload(r)
	if err != nil {
		log.Println(getLogMessage(err, r.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	seatingResponse, err := c.Svc.StartSeating(startSeatingPayload)
	if err != nil {
		log.Println(getLogMessage(err, r.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, seatingResponse) {
		log.Println("error while writing")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) GetTaskResults(w http.ResponseWriter, r *http.Request) {
	taskUuid, err := extractTaskUuid(r)
	if err != nil {
		log.Println(getLogMessage(err, r.URL.Path))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskResults, err := c.Svc.GetTaskResults(taskUuid)
	if err != nil {
		log.Println(getLogMessage(err, r.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, taskResults) {
		log.Println("error while writing")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
