package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	writerErr = "writer error"
)

type Controller struct {
	Svc Service
}

func NewController(svc Service) *Controller {
	return &Controller{Svc: svc}
}

func (c *Controller) GetHallLayout(w http.ResponseWriter, r *http.Request) {
	const funcName = "controller.GetHallLayout"

	hallLayout, err := c.Svc.GetHallLayout()
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, hallLayout) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) StartSeating(w http.ResponseWriter, r *http.Request) {
	const funcName = "controller.StartSeating"

	startSeatingPayload, err := extractStartSeatingPayload(r)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	seatingResponse, err := c.Svc.StartSeating(startSeatingPayload)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, seatingResponse) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) GetSeatingResults(w http.ResponseWriter, r *http.Request) {
	const funcName = "controller.GetSeatingResults"

	taskID, err := extractTaskID(r)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	seatingResults, err := c.Svc.GetSeatingResults(taskID)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, seatingResults) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
