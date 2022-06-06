package hospital

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// PutHospital updates a hospital and returns him/her.
// swagger:route PUT /api/hospitals/:hospitalID Hospital PutHospitalRequest
//
// Update a hospital.
//
// Responses:
//   200: HospitalResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) PutHospital(c *gin.Context) {
	hospitalID := c.Param("hospitalID")
	if hospitalID == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("hospital id is empty")))
		return
	}

	var model models.HospitalModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(apierrors.NewBadRequestError(err))
		return
	}
	// update hospital
	attributes := models.ConvertHospitalModelToAttributes(&model)
	err = u.vaccineClient.CreateOrUpdateHospital(hospitalID, attributes)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, model)
}

// PutHospitalRequest is the request of PutHospital
//
// swagger:parameters PutHospitalRequest
type PutHospitalRequest struct {
	// The hospital info
	// in: body
	Hospital *models.HospitalModel `json:"hospital"`
}
