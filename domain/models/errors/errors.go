package models

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	Business = iota
	Server
)

type Error struct {
	InnerErrors []error
	Location    string
	Level       int
	Title       string
}

func NewServerError() *Error {
	return &Error{Level: Server}
}

func NewBusinessError() *Error {
	return &Error{Level: Business}
}

func (myerr *Error) Error() string {
	var sb strings.Builder
	for i := range myerr.InnerErrors {
		sb.WriteString(myerr.InnerErrors[i].Error())
	}

	errStr := sb.String()

	return fmt.Sprintf("Title: [%s] Errors: [%s] Location: [%s]", myerr.Title, errStr, myerr.Location)
}

func (myerr *Error) Collect(err error, title string) *Error {
	if myerr.Title == "" {
		myerr.Title = title
		myerr.Location = errorLocation()
	}

	myerr.InnerErrors = append(myerr.InnerErrors, err)
	return myerr
}

func (myerr *Error) CollectWithLocation(err error, title string, location string) *Error {
	if myerr.Title == "" {
		myerr.Title = title
		myerr.Location = location
	}

	myerr.InnerErrors = append(myerr.InnerErrors, err)
	return myerr
}

func (myerr *Error) IsEmpty() bool {
	return len(myerr.InnerErrors) == 0
}

func errorLocation() string {
	function, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("File: %s Function: %s Line: %d", chopPath(file), runtime.FuncForPC(function).Name(), line)
}

func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
