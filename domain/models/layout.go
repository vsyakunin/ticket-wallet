package models

type HallLayout struct {
	Name     string    `json:"name"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name string `json:"name"`
	Rows []Row  `json:"rows"`
}

type Row struct {
	Num   int    `json:"num"`
	Seats []Seat `json:"seats"`
}

type Seat struct {
	Num       int    `json:"num"`
	ActualNum int    `json:"actual_num"`
	Rank      string `json:"rank"`
	IsFree    bool   `json:"is_free"`
	IsBlocked bool `json:"is_blocked"`
}
