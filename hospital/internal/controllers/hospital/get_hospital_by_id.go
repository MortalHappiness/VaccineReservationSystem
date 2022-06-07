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
// Parameters:
//   + name: county
//     in: query
//     description: the county of the hospital
//     required: true
//     type: string
//   + name: township
//     in: query
//     description: the township of the hospital
//     required: true
//     type: string
//
// Responses:
//   200: HospitalResponse
//	 400: BadRequestErrorResponse
//   404: NotFoundErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) GetHospitalByID(c *gin.Context) {
	county := c.Query("county")
	if county == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("county is empty")))
		return
	}
	township := c.Query("township")
	if township == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("township is empty")))
		return
	}
	id := c.Param("hospitalID")
	// get hospital info
	row, err := u.vaccineClient.GetHospital(id, county, township)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get hospital: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("hospital#%s#%s#%s not found", county, township, id)))
		return
	}
	hospital, err := models.ConvertRowToHospitalModel(row.Key(), row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to hospital: %w", err)))
		return
	}
	c.JSON(http.StatusOK, hospital)
}
