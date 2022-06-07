package hospital

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/gin-gonic/gin"
)

// DeleteHospital returns hospital id.
// swagger:route DELETE /api/hospitals/:hospitalID Hospital DeleteHospital
//
// Delete the hospital by id.
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
//   200: DeleteHospital hospitalID
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) DeleteHospital(c *gin.Context) {
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
	hospitalID := c.Param("hospitalID")
	if hospitalID == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("hospital id is empty")))
		return
	}

	// delete hospital
	err := u.vaccineClient.DeleteHospital(hospitalID, county, township)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, hospitalID)
}
