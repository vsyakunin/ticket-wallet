package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"ticket-wallet/domain/models"
	"time"
)

type ToBeSeated struct {
	NumPpl    int
	Name      string
	AreSeated bool
}

func (svc *Service) assignSeats(startSeatingPayload models.StartSeatingPayload, taskUuid string) error {
	var seatingResponse models.SeatingResponse

	fileName := fmt.Sprintf(fileNameRaw, folderName, taskUuid)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &seatingResponse)
	if err != nil {
		return err
	}

	seats, err := svc.GetHallLayout()
	if err != nil {
		return err
	}

	seatingResponse.Status = models.SrsProcessing
	err = updateTaskResults(taskUuid, seatingResponse)
	if err != nil {
		return err
	}

	log.Printf("started task UUID %s", taskUuid)

	time.Sleep(30 * time.Second)

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

	err = updateTaskResults(taskUuid, seatingResponse)
	if err != nil {
		return err
	}

	log.Printf("completed task UUID %s", taskUuid)

	return nil
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
