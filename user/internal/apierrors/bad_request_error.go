package apierrors

import "net/http"

// NewBadRequestError returns a new bad request error.
func NewBadRequestError(err error) *APIError {
	return NewAPIError(http.StatusBadRequest, err)
}

// BadRequestErrorResponse is the response of bad request error.
//
// swagger:response BadRequestErrorResponse
type BadRequestErrorResponse struct {
	// in: body
	// required: true
	Error *APIError `json:"error"`
}
