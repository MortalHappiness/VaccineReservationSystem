package apierrors

import (
	"net/http"
	"time"
)

// APIError is the error response for the API.
//
// swagger:model APIError
type APIError struct {
	// The timestamp of the error.
	// in: body
	// required: true
	Timestamp string `json:"timestamp"`
	// The HTTP status code.
	// in: body
	// required: true
	Status int `json:"status"`
	// The status error message.
	// in: body
	// required: true
	ErrMsg string `json:"error"`
	// The error message.
	// in: body
	// required: true
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

// NewAPIError returns a new APIError.
func NewAPIError(status int, err error) *APIError {
	return &APIError{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    status,
		ErrMsg:    http.StatusText(status),
		Message:   err.Error(),
	}
}

// NewBadRequestError returns a new bad request error.
func NewBadRequestError(err error) *APIError {
	return NewAPIError(http.StatusBadRequest, err)
}

// NewInternalServerError returns a new internal server error.
func NewInternalServerError(err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, err)
}

// NewNotFoundError returns a new not found error.
func NewNotFoundError(err error) *APIError {
	return NewAPIError(http.StatusNotFound, err)
}

// NewUnauthorizedError returns a new unauthorized error.
func NewUnauthorizedError(err error) *APIError {
	return NewAPIError(http.StatusUnauthorized, err)
}
