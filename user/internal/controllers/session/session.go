package session

import (
	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetSession(c *gin.Context)
	PostSession(c *gin.Context)
	DelSession(c *gin.Context)
}

// Session handles all
type Session struct {
	env           env.Environments
	vaccineClient *vaccineclient.VaccineClient
}

// Options provides interface to change behavior of User.
type Options struct {
	Env           env.Environments
	VaccineClient *vaccineclient.VaccineClient
}

// New returns default instance of Session.
func New(opt Options) *Session {
	return &Session{
		env:           opt.Env,
		vaccineClient: opt.VaccineClient,
	}
}

// SessionResponse is the response of GetSession/PostSession.
//
// swaggger: response SessionResponse
type SessionResponse struct {
	// The session info
	// in: body
	User *models.UserModel `json:"user"`
}
