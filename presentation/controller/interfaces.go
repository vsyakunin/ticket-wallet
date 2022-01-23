package controller

import (
	"ticket-wallet/domain/models"
)

type Service interface {
	GetHallLayout() (models.HallLayout, error)
}
