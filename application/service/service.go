package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vsyakunin/ticket-wallet/domain/models"
	myerrs "github.com/vsyakunin/ticket-wallet/domain/models/errors"

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

	validationErr = "validation error"
	internalErr   = "internal error"
)

var (
	hallLayoutErr    = errors.New("error while getting hall layout")
	seatingResultErr = errors.New("error while receiving task results")
	startSeatingErr  = errors.New("error while starting seating algorithm")
)

func (svc *Service) GetHallLayout() (models.HallLayout, error) {
	const funcName = "service.GetHallLayout"

	var hallLayout models.HallLayout

	jsonFile, err := os.Open("layout.json")
	if err != nil {
		log.Errorf("%s: error while opening file", funcName)
		return hallLayout, myerrs.NewServerError(internalErr, hallLayoutErr)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Errorf("%s: error while reading file", funcName)
		return hallLayout, myerrs.NewServerError(internalErr, hallLayoutErr)
	}

	err = json.Unmarshal(byteValue, &hallLayout)
	if err != nil {
		log.Errorf("%s: error while unmarshaling file contents", funcName)
		return hallLayout, myerrs.NewServerError(internalErr, hallLayoutErr)
	}

	return hallLayout, nil
}

func (svc *Service) StartSeating(startSeatingReq *models.StartSeatingRequest) (models.SeatingResponse, error) {
	const funcName = "service.StartSeating"

	var seatingResponse models.SeatingResponse

	if err := validateStartSeatingRequest(startSeatingReq); err != nil {
		log.Infof("%s: validation error %v", funcName, err)
		return seatingResponse, err
	}

	taskID := uuid.NewV4().String()

	seatingResponse.TaskID = taskID
	seatingResponse.Status = models.SrsCreated

	newPath := filepath.Join(".", folderName)

	if err := os.MkdirAll(newPath, os.ModePerm); err != nil {
		log.Errorf("%s: error while creating a directory", funcName)
		return seatingResponse, myerrs.NewServerError(internalErr, startSeatingErr)
	}

	fileName := fmt.Sprintf(fileNameRaw, folderName, taskID)

	file, err := json.MarshalIndent(seatingResponse, "", " ")
	if err != nil {
		log.Errorf("%s: error while marshaling json", funcName)
		return seatingResponse, myerrs.NewServerError(internalErr, startSeatingErr)
	}

	if err = ioutil.WriteFile(fileName, file, 0644); err != nil {
		log.Errorf("%s: error while writing to file", funcName)
		return seatingResponse, myerrs.NewServerError(internalErr, startSeatingErr)
	}

	go func() {
		svc.assignSeats(*startSeatingReq, taskID)
	}()

	return seatingResponse, nil
}

func (svc *Service) GetSeatingResults(taskID *string) (models.SeatingResponse, error) {
	const funcName = "service.GetSeatingResults"

	var seatingResponse models.SeatingResponse

	if err := validateGuid(taskID); err != nil {
		return seatingResponse, err
	}

	fileName := fmt.Sprintf(fileNameRaw, folderName, *taskID)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("%s: error while reading file for task ID %s error: %v", funcName, *taskID, err)
		return seatingResponse, myerrs.NewServerError(internalErr, seatingResultErr)
	}

	err = json.Unmarshal(file, &seatingResponse)
	if err != nil {
		log.Errorf("%s: error while unmarshaling file contents for task ID %s error: %v", funcName, *taskID, err)
		return seatingResponse, myerrs.NewServerError(internalErr, seatingResultErr)
	}

	return seatingResponse, nil
}

func updateTaskResults(taskID string, seatingResponse models.SeatingResponse) {
	const funcName = "service.updateTaskResults"

	fileName := fmt.Sprintf(fileNameRaw, folderName, taskID)

	file, err := json.MarshalIndent(seatingResponse, "", " ")
	if err != nil {
		log.Errorf("%s: task ID %s marshaling error: %v", funcName, taskID, err)
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		log.Errorf("%s: task ID %s write to file error: %v", funcName, taskID, err)
	}
}
