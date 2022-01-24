package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	myerrs "github.com/vsyakunin/ticket-wallet/domain/models/errors"

	log "github.com/sirupsen/logrus"
)

const (
	jsonEncodeErr = "error while encoding JSON"
)

func writeJSONResponse(w http.ResponseWriter, path string, records interface{}) bool {
	if records == nil {
		w.WriteHeader(http.StatusOK)
		return true
	}

	bytes, err := json.Marshal(records)
	if err != nil {
		log.Error(getLogMessage(err, path))
		return false
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(bytes); err != nil {
		log.Error(getLogMessage(err, path))
		return false
	}

	return true
}

func writeErrorResponse(w http.ResponseWriter, errsResponse error, instance string) {
	errStr, err := json.Marshal(errsResponse)
	if err != nil {
		log.Error(getLogMessage(err, instance))
		http.Error(w, jsonEncodeErr, http.StatusInternalServerError)
		return
	}

	var statusCode int
	typedErr, ok := errsResponse.(*myerrs.Error)
	if ok {
		switch typedErr.Level {
		case myerrs.Business:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusInternalServerError
	}


	http.Error(w, string(errStr), statusCode)
}

func getLogMessage(err error, instance string) string {
	return fmt.Sprintf("Instance: %s, Error: %s", instance, err.Error())
}