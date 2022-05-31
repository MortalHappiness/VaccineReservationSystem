package hospital

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHospitalV1 returns the hospital information.
// swagger:route GET /v1/hospitals Hospital GetHospital
//
// Get the hospital information.
//
// Responses:
//   200: HospitalResponse
//   500: InternalServerErrorResponse
//
func (u *Hospital) GetHospitalV1(c *gin.Context) {
	var vaccineCnt map[string]int = make(map[string]int)
	vaccineCnt["BNT"] = 1
	model := HospitalModel{
		Name:       "Taipei City Hospital Heping Fuyou Branch",
		ID:         "0001",
		County:     "Taipei",
		Township:   "Zhongzheng District",
		Address:    "No.33, Sec. 2, Zhonghua Rd., Zhongzheng Dist., Taipei City 100058, Taiwan (R.O.C.)",
		VaccineCnt: vaccineCnt,
	}
	c.JSON(http.StatusOK, model)
}
