package apierrors

import "net/http"

// NewNotFoundError returns a new not found error.
func NewNotFoundError(err error) *APIError {
	return NewAPIError(http.StatusNotFound, err)
}

// NotFoundErrorResponse is the response of not found error.
//
// swagger:response NotFoundErrorResponse
type NotFoundErrorResponse struct {
	// in: body
	// required: true
	Error *APIError `json:"error"`
}
