package reservation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteReservationV1 returns reservation id.
// swagger:route DELETE /v1/reservations/:id Reservation DeleteReservation
//
// Delete the reservation by id.
//
// Responses:
//   200: reservation id
//   500: InternalServerErrorResponse
//
func (u *Reservation) DeleteReservationV1(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("id"))
}
