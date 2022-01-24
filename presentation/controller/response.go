package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func getLogMessage(err error, instance string) string {
	return fmt.Sprintf("Instance: %s, Error: %s", instance, err.Error())
}

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
