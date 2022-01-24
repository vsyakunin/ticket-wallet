package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ticket-wallet/domain/models"
)

const (
	taskUuidParam    = "taskUuid"
	paramNotFoundErr = "url parameter %s not found"
)

func extractStartSeatingPayload(r *http.Request) (models.StartSeatingPayload, error) {
	var startSeatingPayload models.StartSeatingPayload
	err := json.NewDecoder(r.Body).Decode(&startSeatingPayload)
	if err != nil {
		return startSeatingPayload, err
	}

	return startSeatingPayload, nil
}

func extractTaskUuid(r *http.Request) (*string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	taskUuid := r.Form.Get(taskUuidParam)
	if taskUuid == "" {
		return nil, fmt.Errorf(paramNotFoundErr, taskUuidParam)
	}

	return &taskUuid, nil
}
