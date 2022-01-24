package controller

import (
	"ticket-wallet/domain/models"
)

type Service interface {
	GetHallLayout() (models.HallLayout, error)
	StartSeating(models.StartSeatingPayload) (models.SeatingResponse, error)
	GetTaskResults(taskID *string) (models.SeatingResponse, error)
}
