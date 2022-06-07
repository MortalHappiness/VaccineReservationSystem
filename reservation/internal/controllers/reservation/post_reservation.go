package reservation

import (
	"fmt"
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

	// add reservation information to bigtable
	attributes := models.ConvertReservationModelToAttributes(&model)
	err = u.vaccineClient.CreateOrUpdateReservation(model.ID, nationID, attributes)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}

	// should modify hositpal left capacity
	// get hospital info
	hospitalRow, err := u.vaccineClient.GetHospital(
		model.Hospital.ID, model.Hospital.County, model.Hospital.Township)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}
	if hospitalRow == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("hospital#%s#%s#%s not found",
			model.Hospital.ID, model.Hospital.County, model.Hospital.Township)))
		return
	}
	hospital, err := models.ConvertRowToHospitalModel(hospitalRow.Key(), hospitalRow)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to hospital: %w", err)))
		return
	}

	if hospital.VaccineCnt[model.VaccineType] <= 0 {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("no vaccine left: %s", model.VaccineType)))
		return
	}
	hospital.VaccineCnt[model.VaccineType]--
	attributes = models.ConvertHospitalModelToAttributes(hospital)
	err = u.vaccineClient.CreateOrUpdateHospital(hospital.ID, hospital.County, hospital.Township, attributes)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}

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
