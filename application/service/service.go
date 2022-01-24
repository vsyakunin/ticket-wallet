package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vsyakunin/ticket-wallet/domain/models"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

const (
	folderName  = "data"
	fileNameRaw = "%s/%s.json"
)

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

	taskUUID := uuid.NewV4().String()

	seatingResponse.TaskID = taskUUID
	seatingResponse.Status = models.SrsCreated

	newPath := filepath.Join(".", folderName)

	if err := os.MkdirAll(newPath, os.ModePerm); err != nil {
		return seatingResponse, err
	}

	fileName := fmt.Sprintf(fileNameRaw, folderName, taskUUID)

	file, err := json.MarshalIndent(seatingResponse, "", " ")
	if err != nil {
		return seatingResponse, err
	}

	if err = ioutil.WriteFile(fileName, file, 0644); err != nil {
		return seatingResponse, err
	}

	go func() {
		svc.assignSeats(startSeatingPayload, taskUUID)
	}()

	return seatingResponse, nil
}

func (svc *Service) GetTaskResults(taskID *string) (models.SeatingResponse, error) {
	var seatingResponse models.SeatingResponse

	fileName := fmt.Sprintf(fileNameRaw, folderName, *taskID)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return seatingResponse, err
	}

	err = json.Unmarshal(file, &seatingResponse)
	if err != nil {
		return seatingResponse, err
	}

	return seatingResponse, nil
}

func updateTaskResults(taskUUID string, seatingResponse models.SeatingResponse) {
	fileName := fmt.Sprintf(fileNameRaw, folderName, taskUUID)

	file, err := json.MarshalIndent(seatingResponse, "", " ")
	if err != nil {
		log.Errorf("can't update file for task UUID %s error: %v", taskUUID, err.Error())
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		log.Errorf("can't update file for task UUID %s error: %v", taskUUID, err.Error())
	}
}
