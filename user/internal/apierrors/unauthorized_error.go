package apierrors

import "net/http"

// NewUnauthorizedError returns a new unauthorized error.
func NewUnauthorizedError(err error) *APIError {
	return NewAPIError(http.StatusUnauthorized, err)
}

// UnauthorizedErrorResponse is the response of unauthorized error.
//
// swagger:response UnauthorizedErrorResponse
type UnauthorizedErrorResponse struct {
	// in: body
	// required: true
	Error *APIError `json:"error"`
}
