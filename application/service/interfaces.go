package service

import (
	"github.com/vsyakunin/ticket-wallet/domain/models"
)

type Repo interface {
	GetHallLayout() (models.HallLayout, error)
	StartSeating(models.StartSeatingRequest) (models.SeatingResponse, error)
	GetSeatingResults(taskID *string) (models.SeatingResponse, error)
}
