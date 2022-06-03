//nolint:revive // use User prefix
package user

import (
	"fmt"
	"strings"

	"cloud.google.com/go/bigtable"
	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
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
	User *UserModel `json:"user"`
}

// UserModel is the body format of UserResponse
//
// swagger:model UserModel
type UserModel struct {
	// The user name
	// example: bob
	// in: body
	// required: false
	Name string `json:"name"`
	// The user gender
	// example: male
	// in: body
	// required: false
	Gender string `json:"gender"`
	// The user nation ID
	// example: A123456789
	// in: body
	// required: true
	NationID string `json:"nationID" binding:"required"`
	// The user healthCardID
	// example: 000011112222
	// in: body
	// required: true
	HealthCardID string `json:"healthCardID"`
	// The user birthday
	// example: 2022/05/23
	// in: body
	// required: false
	BirthDay string `json:"birthDay"`
	// The user address
	// example: No.1, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan
	// in: body
	// required: false
	Address string `json:"address"`
	// The user phone number
	// example: 0912345678
	// in: body
	// required: false
	Phone string `json:"phone"`
	// The user inoculated vaccines
	// example: [AZ, BNT]
	// in: body
	// required: false
	Vaccines []string `json:"vaccines"`
}

// ConvertRowToUserModel converts bigtable.Row to *UserModel with given nationID.
func ConvertRowToUserModel(nationID string, row bigtable.Row) (*UserModel, error) {
	user := &UserModel{
		NationID: nationID,
	}
	for _, col := range row["user"] {
		qualifier := col.Column[strings.IndexByte(col.Column, ':')+1:]
		switch qualifier {
		case "name":
			user.Name = string(col.Value)
		case "gender":
			user.Gender = string(col.Value)
		case "healthCardID":
			user.HealthCardID = string(col.Value)
		case "birthday":
			user.BirthDay = string(col.Value)
		case "address":
			user.Address = string(col.Value)
		case "phone":
			user.Phone = string(col.Value)
		case "vaccines":
			user.Vaccines = strings.Split(string(col.Value), ",")
		default:
			return nil, fmt.Errorf("unknown qualifier: %s", qualifier)
		}
	}
	return user, nil
}

func (u *User) AuthVerify(c *gin.Context, givenNationID string) error {
	nationID, exists := c.Get("nationID")
	if !exists {
		return fmt.Errorf("nationID not found in context")
	}
	if givenNationID != nationID {
		return fmt.Errorf("nationID not match")
	}
	return nil
}
