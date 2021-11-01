package errors

import "net/http"

type RestError struct {
	Message string `json: "message"`
	Status  int    `json: "status"`
	Error   string `json: "error"`
}

func NewBadRequest(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFound(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}