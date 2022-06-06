package reservation

import (
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/gin-gonic/gin"
)

// DeleteReservationV1 returns reservation id.
// swagger:route DELETE /api/reservations/users/:nationID/:reservationID Reservation DeleteReservation
//
// Delete the reservation by id.
//
// Responses:
//   200: DeleteReservationResponse reservationID
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) DeleteReservation(c *gin.Context) {
	nationID := c.Param("nationID")
	err := AuthVerify(c, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
		return
	}

	reservationID := c.Param("reservationID")
	// delete reservation information from bigtable'
	err = u.vaccineClient.DeleteReservation(reservationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, reservationID)
}
