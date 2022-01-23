package service

import (
	"ticket-wallet/domain/models"
)

type ToBeSeated struct {
	NumPpl    int
	Name      string
	AreSeated bool
}

func assignSeats(seats models.HallLayout, startSeatingPayload models.StartSeatingPayload) models.HallLayout {
	toBeSeated := make([]ToBeSeated, 0, len(startSeatingPayload.Groups))
	for _, group := range startSeatingPayload.Groups {
		toBeSeated = append(toBeSeated, ToBeSeated{NumPpl: group.GroupSize, Name: group.Name, AreSeated: false})
	}

	// i is a number of group of groups
	for i := range toBeSeated {
		seatGroup(seats, toBeSeated[i].NumPpl, toBeSeated[i].Name)

		toBeSeated[i].AreSeated = true
	}

	return seats
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
