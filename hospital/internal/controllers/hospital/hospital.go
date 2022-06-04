//nolint:revive // use Hospital prefix
package hospital

import (
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/MortalHappiness/VaccineReservationSystem/hospital/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetHospital(c *gin.Context)
	PostHospital(c *gin.Context)
	PutHospital(c *gin.Context)
	DeleteHospital(c *gin.Context)
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

// HospitalResponse is the response of GetHospital
//
// swagger:response HospitalResponse
type HospitalResponse struct {
	// The hospital info
	// in: body
	Hospital *models.HospitalModel `json:"hospital"`
}
