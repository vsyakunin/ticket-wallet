package models

type HallLayout struct {
	Name     string    `json:"name,omitempty"`
	Sections []Section `json:"sections,omitempty"`
}

type Section struct {
	Name     string `json:"name"`
	IsCurved bool   `json:"is_curved"`
	Rows     []Row  `json:"rows"`
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
	IsBlocked bool   `json:"is_blocked"`
	TakenBy   string `json:"taken_by"`
}
