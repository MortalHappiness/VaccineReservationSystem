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
//
// Responses:
//   200: DeleteHospital hospitalID
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) DeleteHospital(c *gin.Context) {
	hospitalID := c.Param("hospitalID")
	if hospitalID == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("hospital id is empty")))
		return
	}

	// TODO: delete hospital

	c.JSON(http.StatusOK, hospitalID)
}
