package apierrors

import "net/http"

// NewInternalServerError returns a new internal server error.
func NewInternalServerError(err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, err)
}

// InternalServerErrorResponse is the response of internal server error.
//
// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	// in: body
	// required: true
	Error *APIError `json:"error"`
}
