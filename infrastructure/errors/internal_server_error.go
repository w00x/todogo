package errors

import "net/http"

type InternalServerError struct {
	BaseError
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{BaseError{message, "INTERNAL_SERVER_ERROR", 1001, http.StatusInternalServerError}}
}