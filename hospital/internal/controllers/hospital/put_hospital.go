package hospital

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PutHospitalV1 updates a hospital and returns him/her.
// swagger:route PUT /v1/hospitals/:id Hospital PutHospitalRequest
//
// Update a hospital.
//
// Responses:
//   200: HospitalResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) PutHospitalV1(c *gin.Context) {
	var model HospitalModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(fmt.Errorf("put hospital request: %w", err))
		return
	}
	// id := c.Param("id")
	// TODO check hospital exists
	// TODO update
	c.JSON(http.StatusOK, model)
}

// PutHospitalRequest is the request of PutHospitalV1
//
// swagger:parameters PutHospitalRequest
type PutHospitalRequest struct {
	// The hospital info
	// in: body
	Hospital *HospitalModel `json:"hospital"`
}
