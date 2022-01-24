package service

import (
	"errors"
	"fmt"
	"github.com/vsyakunin/ticket-wallet/domain/models"

	myerrs "github.com/vsyakunin/ticket-wallet/domain/models/errors"

	uuid "github.com/satori/go.uuid"
)

const (
	invalidParameterErr = "invalid parameter"
)

func validateGuid(taskID *string) error {
	if taskID == nil {
		err := errors.New("incorrect task ID")
		return myerrs.NewBusinessError(validationErr, err)
	}

	if _, err := uuid.FromString(*taskID); err != nil {
		return myerrs.NewBusinessError(validationErr, err)
	}

	return nil
}

func validateStartSeatingRequest(params *models.StartSeatingRequest) error {
	for i, group := range params.Groups {
		if group.GroupSize < 1 {
			err := errors.New(fmt.Sprintf("group #%d: size can't be smaller than 1", i + 1))
			return myerrs.NewBusinessError(invalidParameterErr, err)
		}

		if group.Name == "" {
			err := errors.New(fmt.Sprintf("group #%d: name can't be empty", i + 1))
			return myerrs.NewBusinessError(invalidParameterErr, err)
		}
	}

	return nil
}
