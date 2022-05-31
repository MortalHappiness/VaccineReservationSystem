package reservation

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostReservationV1 adds a new reservation and returns him/her.
// swagger:route POST /v1/reservations Reservation PostReservationRequest
//
// Add a new reservation.
//
// Responses:
//   200: ReservationResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) PostReservationV1(c *gin.Context) {
	var model ReservationModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(fmt.Errorf("post reservation request: %w", err))
		return
	}
	c.JSON(http.StatusOK, model)
}

// PostReservationRequest is the request of PostReservationV1
//
// swagger:parameters PostReservationRequest
type PostReservationRequest struct {
	// The reservation info
	// in: body
	Reservation *ReservationModel `json:"reservation"`
}
