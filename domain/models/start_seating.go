package models

type StartSeatingRequest struct {
	Groups []SeatingGroup `json:"groups"`
}

type SeatingGroup struct {
	GroupSize int    `json:"size"`
	Name      string `json:"name"`
}
