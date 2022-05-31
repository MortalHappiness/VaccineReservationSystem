package hospital

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteHospitalV1 returns hospital id.
// swagger:route DELETE /v1/hospitals/:id Hospital DeleteHospital
//
// Delete the hospital by id.
//
// Responses:
//   200: hospital id
//   500: InternalServerErrorResponse
//
func (u *Hospital) DeleteHospitalV1(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("id"))
}
