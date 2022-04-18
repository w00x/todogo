package errors

import "net/http"

type NotCreatedError struct {
	BaseError
}

func NewNotCreatedError(message string) *NotCreatedError {
	return &NotCreatedError{BaseError{message, "NOT_FOUND", 1002, http.StatusBadRequest}}
}