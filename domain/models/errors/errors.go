package models

import (
	"fmt"
)

const (
	Business = iota
	Server
)

type Error struct {
	Level       int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewServerError(title string, err error) *Error {
	return &Error{Level: Server, Title: title, Description: err.Error()}
}

func NewBusinessError(title string, err error) *Error {
	return &Error{Level: Business, Title: title, Description: err.Error()}
}

func (myerr *Error) Error() string {
	return fmt.Sprintf("Title: [%s] Description: [%s]", myerr.Title, myerr.Description)
}
