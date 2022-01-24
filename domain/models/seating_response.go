package models

type SeatingResponseStatus string

const (
	SrsCreated    SeatingResponseStatus = "created"
	SrsProcessing SeatingResponseStatus = "processing"
	SrsCompleted  SeatingResponseStatus = "completed"
	SrsError      SeatingResponseStatus = "error"
)

type SeatingResponse struct {
	TaskID  string                `json:"task_id"`
	Status  SeatingResponseStatus `json:"status"`
	Payload HallLayout            `json:"payload,omitempty"`
}
