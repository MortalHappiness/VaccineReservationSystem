package infoexample

import (
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetInfoV1(c *gin.Context)
	PostInfoV1(c *gin.Context)
}

// InfoExample handles all info-related requests.
type InfoExample struct {
	env env.Environments
}

// Options provides interface to change behavior of InfoExample.
type Options struct {
	Env env.Environments
}

// New returns default instance of InfoExample.
func New(opt Options) *InfoExample {
	return &InfoExample{
		env: opt.Env,
	}
}

// InfoModel is the body format of PostInfoV1
// Info Model
//
// swagger:model infoModel
type InfoModel struct {
	// The Name info
	// example: Trump
	// required: true
	Name string `json:"name"`
}
