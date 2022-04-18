package errors

import "net/http"

type NotDeletedError struct {
	BaseError
}

func NewNotDeletedError(message string) *NotDeletedError {
	return &NotDeletedError{BaseError{message, "NOT_FOUND", 1003, http.StatusBadRequest}}
}