package hospital

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// GetHospital returns the hospital information.
// swagger:route GET /api/hospitals Hospital GetHospital
//
// Get the hospital information.
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
func (u *Hospital) GetHospital(c *gin.Context) {
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

	// TODO: Get the hospital information
	model := models.HospitalModel{
		Name:     "Taipei City Hospital Heping Fuyou Branch",
		ID:       "0001",
		County:   "臺北市",
		Township: "中正區",
		Address:  "No.33, Sec. 2, Zhonghua Rd., Zhongzheng Dist., Taipei City 100058, Taiwan (R.O.C.)",
		VaccineCnt: map[string]int{
			"BNT": 100,
			"AZ":  200,
		},
	}

	c.JSON(http.StatusOK, model)
}
