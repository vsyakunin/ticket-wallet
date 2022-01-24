package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vsyakunin/ticket-wallet/domain/models"
	myerrs "github.com/vsyakunin/ticket-wallet/domain/models/errors"
)

const (
	taskIDParam = "taskId"

	formParseErr     = "couldn't fill the form with parameters"
	paramNotFoundErr = "url parameter %s not found"
	paramParseErr    = "parameter parsing error"
	jsonParseErr     = "can't read JSON request body"
)

func extractStartSeatingPayload(r *http.Request) (*models.StartSeatingRequest, error) {
	var startSeatingPayload models.StartSeatingRequest
	err := json.NewDecoder(r.Body).Decode(&startSeatingPayload)
	if err != nil {
		return nil, myerrs.NewBusinessError(jsonParseErr, err)
	}

	return &startSeatingPayload, nil
}

func extractTaskID(r *http.Request) (*string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, myerrs.NewBusinessError(formParseErr, err)
	}

	taskID := r.Form.Get(taskIDParam)
	if taskID == "" {
		err := fmt.Errorf(paramNotFoundErr, taskIDParam)
		return nil, myerrs.NewBusinessError(paramParseErr, err)
	}

	return &taskID, nil
}
