package controller

import (
	"encoding/json"
	"net/http"

	"ticket-wallet/domain/models"
)

func extractStartSeatingPayload(r *http.Request) (models.StartSeatingPayload, error) {
	var startSeatingPayload models.StartSeatingPayload
	err := json.NewDecoder(r.Body).Decode(&startSeatingPayload)
	if err != nil {
		return startSeatingPayload, err
	}

	return startSeatingPayload, nil
}
