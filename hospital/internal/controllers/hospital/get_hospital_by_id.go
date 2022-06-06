package hospital

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// GetHospitalByID returns the hospital information.
// swagger:route GET /api/hospitals/:hospitalID Hospital GetHospitalByID
//
// Get the hospital information by id.
//
// Responses:
//   200: HospitalResponse
//	 400: BadRequestErrorResponse
//   404: NotFoundErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) GetHospitalByID(c *gin.Context) {
	id := c.Param("hospitalID")
	// get hospital info
	row, err := u.vaccineClient.GetHospital(id)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get hospital: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("hospital %s not found", id)))
		return
	}
	hospital, err := models.ConvertRowToHospitalModel(id, row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to hospital: %w", err)))
		return
	}
	c.JSON(http.StatusOK, hospital)
}
