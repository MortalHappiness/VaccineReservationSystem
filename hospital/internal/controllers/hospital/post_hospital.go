package hospital

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostHospitalV1 adds a new hospital and returns him/her.
// swagger:route POST /v1/hospitals Hospital PostHospitalRequest
//
// Add a new hospital.
//
// Responses:
//   200: HospitalResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) PostHospitalV1(c *gin.Context) {
	var model HospitalModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(fmt.Errorf("post hospital request: %w", err))
		return
	}
	c.JSON(http.StatusOK, model)
}

// PostHospitalRequest is the request of PostHospitalV1
//
// swagger:parameters PostHospitalRequest
type PostHospitalRequest struct {
	// The hospital info
	// in: body
	Hospital *HospitalModel `json:"hospital"`
}
