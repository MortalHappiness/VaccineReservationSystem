//nolint:revive // use Hospital prefix
package hospital

import (
	"github.com/MortalHappiness/VaccineReservationSystem/hospital/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetHospitalV1(c *gin.Context)
	PostHospitalV1(c *gin.Context)
	// PutHospitalV1(c *gin.Context)
	DeleteHospitalV1(c *gin.Context)
}

// Hospital handles all info-related requests.
type Hospital struct {
	env env.Environments
}

// Options provides interface to change behavior of Hospital.
type Options struct {
	Env env.Environments
}

// New returns default instance of Hospital.
func New(opt Options) *Hospital {
	return &Hospital{
		env: opt.Env,
	}
}

// HospitalResponse is the response of GetHospitalV1
//
// swagger:response HospitalResponse
type HospitalResponse struct {
	// The hospital info
	// in: body
	Hospital *HospitalModel `json:"hospital"`
}

// HospitalModel is the body format of HospitalResponse
//
// swagger:model HospitalModel
type HospitalModel struct {
	// The hospital name
	// required: true
	// example: Taipei City Hospital Heping Fuyou Branch
	Name string `json:"name"`
	// The hospital ID
	// example: 0001
	// required: true
	ID string `json:"id"`
	// The hospital county
	// example: Taipei
	// required: true
	County string `json:"county"`
	// The hospital township
	// example: Zhongzheng District
	// required: true
	Township string `json:"township"`
	// The hospital address
	// example: No.33, Sec. 2, Zhonghua Rd., Zhongzheng Dist., Taipei City 100058, Taiwan (R.O.C.)
	// required: true
	Address string `json:"address"`
	// The hospital vaccine
	// example: {"BNT":1}
	// required: true
	VaccineCnt map[string]int `json:"vaccinecnt"`
}
