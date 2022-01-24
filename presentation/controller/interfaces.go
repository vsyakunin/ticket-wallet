package controller

import (
	"github.com/vsyakunin/ticket-wallet/domain/models"
)

type Service interface {
	GetHallLayout() (models.HallLayout, error)
	StartSeating(*models.StartSeatingRequest) (models.SeatingResponse, error)
	GetSeatingResults(taskID *string) (models.SeatingResponse, error)
}
