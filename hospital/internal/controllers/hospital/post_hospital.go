package hospital

import (
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// PostHospital adds a new hospital and returns him/her.
// swagger:route POST /api/hospitals Hospital PostHospitalRequest
//
// Add a new hospital.
//
// Responses:
//   200: HospitalResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) PostHospital(c *gin.Context) {
	var model models.HospitalModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(apierrors.NewBadRequestError(err))
		return
	}

	// add hospital
	attributes := models.ConvertHospitalModelToAttributes(&model)
	err = u.vaccineClient.CreateOrUpdateHospital(model.ID, model.County, model.Township, attributes)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, model)
}

// PostHospitalRequest is the request of PostHospital
//
// swagger:parameters PostHospitalRequest
type PostHospitalRequest struct {
	// The hospital info
	// in: body
	Hospital *models.HospitalModel `json:"hospital"`
}
