//nolint:revive // use User prefix
package user

import (
	"fmt"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of user.
type I interface {
	GetUserByID(c *gin.Context)
	PostUser(c *gin.Context)
	PutUser(c *gin.Context)
	DelUser(c *gin.Context)
}

// User handles all info-related requests.
type User struct {
	env           env.Environments
	vaccineClient *vaccineclient.VaccineClient
}

// Options provides interface to change behavior of User.
type Options struct {
	Env           env.Environments
	VaccineClient *vaccineclient.VaccineClient
}

// New returns default instance of User.
func New(opt Options) *User {
	return &User{
		env:           opt.Env,
		vaccineClient: opt.VaccineClient,
	}
}

// UserResponse is the response of GetUser/PostUser/PutUser.
//
// swagger:response UserResponse
type UserResponse struct {
	// The user info
	// in: body
	User *models.UserModel `json:"user"`
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
