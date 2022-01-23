package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Println(err.Error())
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
	fmt.Println(startSeatingPayload.Groups)

	var seatingResponse models.SeatingResponse

	taskID, err := uuid.NewV4()
	if err != nil {
		return seatingResponse, err
	}

	return models.SeatingResponse{
		TaskID: taskID.String(),
		Status: models.SrsProcessing,
	}, nil
}
