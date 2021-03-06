package reservation

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// GetReservationByID returns the reservation information.
// swagger:route GET /api/reservations/users/:nationID/:reservationID Reservation GetReservation
//
// Get the reservation information by reservation id.
//
// Responses:
//   200: ReservationResponse
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   404: NotFoundErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Reservation) GetReservationByID(c *gin.Context) {
	nationID := c.Param("nationID")
	err := AuthVerify(c, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
		return
	}

	reservationID := c.Param("reservationID")

	// get reservation info
	row, err := u.vaccineClient.GetReservation(reservationID, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get reservation: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("reservation#%s#%s not found", nationID, reservationID)))
		return
	}

	reservation, err := models.ConvertRowToReservationModel(row.Key(), row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to reservation: %w", err)))
		return
	}

	// cannot get others reservation info
	err = AuthVerify(c, reservation.User.NationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
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

	user, err := models.ConvertRowToUserModel(reservation.User.NationID, userRow)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to user: %w", err)))
		return
	}

	reservation.User = user

	// get hospital info
	hospitalRow, err := u.vaccineClient.GetHospital(
		reservation.Hospital.ID, reservation.Hospital.County, reservation.Hospital.Township)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get hospital: %w", err)))
		return
	}
	if hospitalRow == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("hospital#%s#%s#%s not found",
			reservation.Hospital.County, reservation.Hospital.Township, reservation.Hospital.ID)))
		return
	}
	hospital, err := models.ConvertRowToHospitalModel(hospitalRow.Key(), hospitalRow)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to hospital: %w", err)))
		return
	}
	reservation.Hospital = hospital

	c.JSON(http.StatusOK, reservation)
}
