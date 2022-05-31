//nolint:revive // use Reservation prefix
package reservation

import (
	"github.com/MortalHappiness/VaccineReservationSystem/reservation/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetReservationV1(c *gin.Context)
	PostReservationV1(c *gin.Context)
	DeleteReservationV1(c *gin.Context)
}

// Reservation handles all info-related requests.
type Reservation struct {
	env env.Environments
}

// Options provides interface to change behavior of Reservation.
type Options struct {
	Env env.Environments
}

// New returns default instance of Reservation.
func New(opt Options) *Reservation {
	return &Reservation{
		env: opt.Env,
	}
}

// ReservationResponse is the response of GetReservationV1
//
// swagger:response ReservationResponse
type ReservationResponse struct {
	// The reservation info
	// in: body
	Reservation *ReservationModel `json:"reservation"`
}

// ReservationModel is the body format of ReservationResponse
//
// swagger:model ReservationModel
type ReservationModel struct {
	// The reservation id
	// required: true
	// example: 0001
	ID string `json:"id"`
	// The reservation of the user
	// example: Bob
	// required: true
	User string `json:"user"`
	// The reservation hospital
	// example: Taipei City Hospital Heping Fuyou Branch
	// required: true
	Hospital string `json:"hospital"`
	// The reservation vaccinetype
	// example: BNT
	// required: true
	VaccineType string `json:"vaccinetype"`
	// The reservation date
	// example: 1653974953
	// required: true
	Date int64 `json:"date"`
	// The vaccination is completed
	// example: true
	// required: true
	Completed bool `json:"completed"`
}
