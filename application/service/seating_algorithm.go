package service

import (
	"fmt"
	"strconv"

	"ticket-wallet/domain/models"
)

type ToBeSeated struct {
	NumPpl    int
	AreSeated bool
}

func assignSeats(seats models.HallLayout, groups []int) {
	toBeSeated := make([]ToBeSeated, 0, len(groups))
	for _, group := range groups {
		toBeSeated = append(toBeSeated, ToBeSeated{NumPpl: group, AreSeated: false})
	}

	// i is a number of group of groups
	for i := range toBeSeated {
		seatGroup(seats, toBeSeated[i].NumPpl, i)

		toBeSeated[i].AreSeated = true
	}

	fmt.Printf("all seats assigned \n%v\n", seats)
}

func seatGroup(seats models.HallLayout, group, count int) {
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
						seats.Sections[sec].Rows[r].Seats[seatNum].TakenBy = strconv.Itoa(count + 1)
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
