package responses

import "github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"

// BadRequestErrorResponse is the response of bad request error.
//
// swagger:response BadRequestErrorResponse
type BadRequestErrorResponse struct {
	// in: body
	// required: true
	Error *apierrors.APIError `json:"error"`
}

// InternalServerErrorResponse is the response of internal server error.
//
// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	// in: body
	// required: true
	Error *apierrors.APIError `json:"error"`
}

// UnauthorizedErrorResponse is the response of unauthorized error.
//
// swagger:response UnauthorizedErrorResponse
type UnauthorizedErrorResponse struct {
	// in: body
	// required: true
	Error *apierrors.APIError `json:"error"`
}

// NotFoundErrorResponse is the response of not found error.
//
// swagger:response NotFoundErrorResponse
type NotFoundErrorResponse struct {
	// in: body
	// required: true
	Error *apierrors.APIError `json:"error"`
}
