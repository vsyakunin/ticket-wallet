package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/vsyakunin/ticket-wallet/domain/models"

	log "github.com/sirupsen/logrus"
)

type ToBeSeated struct {
	NumPpl    int
	Name      string
	AreSeated bool
}

func (svc *Service) assignSeats(startSeatingPayload models.StartSeatingRequest, taskID string) {
	const funcName = "service.assignSeats"

	var seatingResponse models.SeatingResponse

	fileName := fmt.Sprintf(fileNameRaw, folderName, taskID)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("%s: error while reading file for task UUID %s error: %v", funcName, taskID, err)
		seatingResponse.Status = models.SrsError
		updateTaskResults(taskID, seatingResponse)
		return
	}

	err = json.Unmarshal(file, &seatingResponse)
	if err != nil {
		log.Errorf("%s: error while unmarshaling file contents for task UUID %s error: %v", funcName, taskID, err.Error())
		seatingResponse.Status = models.SrsError
		updateTaskResults(taskID, seatingResponse)
		return
	}

	seats, err := svc.GetHallLayout()
	if err != nil {
		log.Errorf("%s: error while getting hall layout for task UUID %s error: %v", funcName, taskID, err.Error())
		seatingResponse.Status = models.SrsError
		updateTaskResults(taskID, seatingResponse)
		return
	}

	seatingResponse.Status = models.SrsProcessing
	updateTaskResults(taskID, seatingResponse)

	log.Infof("%s: started task UUID %s", funcName, taskID)

	toBeSeated := make([]ToBeSeated, 0, len(startSeatingPayload.Groups))
	for _, group := range startSeatingPayload.Groups {
		toBeSeated = append(toBeSeated, ToBeSeated{NumPpl: group.GroupSize, Name: group.Name, AreSeated: false})
	}

	// i is a number of group of groups
	for i := range toBeSeated {
		seatGroup(seats, toBeSeated[i].NumPpl, toBeSeated[i].Name)

		toBeSeated[i].AreSeated = true
	}

	seatingResponse.Status = models.SrsCompleted
	seatingResponse.Payload = seats

	updateTaskResults(taskID, seatingResponse)

	log.Infof("%s: completed task UUID %s", funcName, taskID)
}

func seatGroup(seats models.HallLayout, group int, name string) {
	for j := 0; j < group; j++ {
		var seatAssigned bool

		// sec is a section number
		for sec := range seats.Sections {

			// r is row number
			for r := range seats.Sections[sec].Rows {
				row := seats.Sections[sec].Rows[r]

				// s is seat number
				for s := range row.Seats {

					var seatNum int
					fromStart := r%2 == 0
					if !fromStart {
						seatNum = len(row.Seats) - 1 - s
					} else {
						seatNum = s
					}

					if seats.Sections[sec].Rows[r].Seats[seatNum].IsFree {
						seats.Sections[sec].Rows[r].Seats[seatNum].TakenBy = name
						seats.Sections[sec].Rows[r].Seats[seatNum].IsFree = false
						seatAssigned = true
						break
					}
				}

				if !seatAssigned {
					continue
				} else {
					break
				}
			}

			if !seatAssigned {
				continue
			} else {
				break
			}
		}

	}
}
