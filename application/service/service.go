package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"ticket-wallet/domain/models"
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
