package reservation

import (
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// PostReservation adds a new reservation and returns him/her.
// swagger:route POST /api/reservations/users/:nationID Reservation PostReservationRequest
//
// Add a new reservation.
//
// Responses:
//   200: ReservationResponse
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) PostReservation(c *gin.Context) {
	nationID := c.Param("nationID")
	err := AuthVerify(c, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
		return
	}

	var model models.ReservationModel
	err = c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(apierrors.NewBadRequestError(err))
		return
	}

	// TODO: add reservation information to bigtable

	c.JSON(http.StatusOK, model)
}

// PostReservationRequest is the request of PostReservation
//
// swagger:parameters PostReservationRequest
type PostReservationRequest struct {
	// The reservation info
	// in: body
	Reservation *models.ReservationModel `json:"reservation"`
}
