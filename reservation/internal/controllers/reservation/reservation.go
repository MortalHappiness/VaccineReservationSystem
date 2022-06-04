//nolint:revive // use Reservation prefix
package reservation

import (
	"fmt"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/MortalHappiness/VaccineReservationSystem/reservation/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetReservation(c *gin.Context)
	PostReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
}

// Reservation handles all info-related requests.
type Reservation struct {
	env           env.Environments
	vaccineClient *vaccineclient.VaccineClient
}

// Options provides interface to change behavior of Reservation.
type Options struct {
	Env           env.Environments
	VaccineClient *vaccineclient.VaccineClient
}

// New returns default instance of Reservation.
func New(opt Options) *Reservation {
	return &Reservation{
		env:           opt.Env,
		vaccineClient: opt.VaccineClient,
	}
}

// ReservationResponse is the response of GetReservation
//
// swagger:response ReservationResponse
type ReservationResponse struct {
	// The reservation info
	// in: body
	Reservation *models.ReservationModel `json:"reservation"`
}

func AuthVerify(c *gin.Context, givenNationID string) error {
	nationID, exists := c.Get("nationID")
	if !exists {
		return fmt.Errorf("nationID not found in context")
	}
	if givenNationID != nationID {
		return fmt.Errorf("nationID not match")
	}
	return nil
}
