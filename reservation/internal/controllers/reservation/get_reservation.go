package reservation

import (
	"fmt"
	"net/http"

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
	// get reservation information from bigtable
	rows, err := u.vaccineClient.GetReservations(nationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get reservation: %w", err)))
		return
	}
	if len(rows) == 0 {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("reservation %s not found", nationID)))
		return
	}

	reservations := []models.ReservationModel{}
	for _, row := range rows {
		reservation, err := models.ConvertRowToReservationModel(row.Key(), row)
		if err != nil {
			_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to reservation: %w", err)))
			return
		}

		// get user info
		userRow, err := u.vaccineClient.GetUser(reservation.User.NationID)
		if err != nil {
			_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get user: %w", err)))
			return
		}
		if userRow == nil {
			_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("user %s not found", reservation.User.NationID)))
			return
		}
		reservation.User, err = models.ConvertRowToUserModel(userRow.Key(), userRow)
		if err != nil {
			_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to user: %w", err)))
			return
		}

		// get hospital info
		hospitalRow, err := u.vaccineClient.GetHospital(
			reservation.Hospital.ID, reservation.Hospital.County, reservation.Hospital.Township)
		if err != nil {
			_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get hospital: %w", err)))
			return
		}
		if hospitalRow == nil {
			_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("hospital %s not found", reservation.Hospital.ID)))
			return
		}
		hospital, err := models.ConvertRowToHospitalModel(reservation.Hospital.ID, hospitalRow)
		if err != nil {
			_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to hospital: %w", err)))
			return
		}
		reservation.Hospital = hospital

		reservations = append(reservations, *reservation)
	}

	c.JSON(http.StatusOK, reservations)
}
