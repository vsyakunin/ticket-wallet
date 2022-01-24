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

	taskUuid, err := extractTaskUuid(r)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	taskResults, err := c.Svc.GetSeatingResults(taskUuid)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, taskResults) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
