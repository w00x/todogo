package errors

import "net/http"

type NotFoundError struct {
	BaseError
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{BaseError{message, "NOT_FOUND", 1000, http.StatusNotFound}}
}