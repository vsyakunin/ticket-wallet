package service

import (
	"ticket-wallet/domain/models"
)

type Repo interface {
	GetHallLayout() (models.HallLayout, error)
}
