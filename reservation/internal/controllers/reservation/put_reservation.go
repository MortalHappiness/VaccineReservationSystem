package reservation

import (
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// Put adds a new reservation and returns him/her.
// swagger:route PUT /api/reservations/users/:nationID/:reservationID Reservation PutReservationRequest
//
// Add a new reservation.
//
// Responses:
//   200: ReservationResponse
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) PutReservation(c *gin.Context) {
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

	// add reservation information to bigtable
	attributes := models.ConvertReservationModelToAttributes(&model)
	err = u.vaccineClient.CreateOrUpdateReservation(model.ID, nationID, attributes)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, model)
}

// PutReservationRequest is the request of PutReservation
//
// swagger:parameters PutReservationRequest
type PutReservationRequest struct {
	// The reservation info
	// in: body
	Reservation *models.ReservationModel `json:"reservation"`
}
