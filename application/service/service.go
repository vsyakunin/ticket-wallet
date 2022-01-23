package service

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"ticket-wallet/domain/models"

	uuid "github.com/satori/go.uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (svc *Service) GetHallLayout() (models.HallLayout, error) {
	var hallLayout models.HallLayout

	jsonFile, err := os.Open("layout.json")
	if err != nil {
		return hallLayout, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return hallLayout, err
	}

	err = json.Unmarshal(byteValue, &hallLayout)
	if err != nil {
		return hallLayout, err
	}

	return hallLayout, nil
}

func (svc *Service) StartSeating(startSeatingPayload models.StartSeatingPayload) (models.SeatingResponse, error) {
	var seatingResponse models.SeatingResponse

	hallLayout, err := svc.GetHallLayout()
	if err != nil {
		return seatingResponse, err
	}

	seatedLayout := assignSeats(hallLayout, startSeatingPayload)

	taskID, err := uuid.NewV4()
	if err != nil {
		return seatingResponse, err
	}

	return models.SeatingResponse{
		TaskID:  taskID.String(),
		Status:  models.SrsProcessing,
		Payload: seatedLayout,
	}, nil
}
