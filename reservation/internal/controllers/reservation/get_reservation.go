package reservation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReservationV1 returns the reservation information.
// swagger:route GET /v1/reservations Reservation GetReservation
//
// Get the reservation information.
//
// Responses:
//   200: ReservationResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) GetReservationV1(c *gin.Context) {
	var vaccineCnt map[string]int = make(map[string]int)
	vaccineCnt["BNT"] = 1
	model := ReservationModel{
		ID:          "0001",
		Hospital:    "Taipei City Reservation Heping Fuyou Branch",
		User:        "Bob",
		VaccineType: "BNT",
		Date:        1653974953,
		Completed:   true,
	}
	c.JSON(http.StatusOK, model)
}
