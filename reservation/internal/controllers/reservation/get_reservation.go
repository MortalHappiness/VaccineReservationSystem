package reservation

import (
	"net/http"
	"time"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// GetReservation returns the reservation information.
// swagger:route GET /api/reservations/users/:nationID Reservation GetReservation
//
// Get the reservation information.
//
// Responses:
//   200: ReservationResponse
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   404: NotFoundErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) GetReservation(c *gin.Context) {
	nationID := c.Param("nationID")
	err := AuthVerify(c, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
		return
	}
	// TODO: get reservation information from bigtable

	// sample reservation information
	model := models.ReservationModel{
		ID: "0001",
		Hospital: &models.HospitalModel{
			ID:   "0001",
			Name: "Hospital 1",
		},
		User: &models.UserModel{
			NationID: nationID,
			Name:     "John Doe",
		},
		VaccineType: "BNT",
		Date:        time.Now().Unix(),
		Completed:   true,
	}

	c.JSON(http.StatusOK, model)
}
