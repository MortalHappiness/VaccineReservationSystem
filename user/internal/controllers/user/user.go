//nolint:revive // use User prefix
package user

import (
	"strings"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/gin-gonic/gin"
)

// I is interface of authentication.
type I interface {
	GetUserV1(c *gin.Context)
	PostUserV1(c *gin.Context)
}

// User handles all info-related requests.
type User struct {
	env env.Environments
}

// Options provides interface to change behavior of User.
type Options struct {
	Env env.Environments
}

// New returns default instance of User.
func New(opt Options) *User {
	return &User{
		env: opt.Env,
	}
}

// UserResponse is the response of GetUserV1
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
	// required: true
	// example: bob
	Name string `json:"name"`
	// The user gender
	// example: male
	// required: true
	Gender string `json:"gender"`
	// The user ID
	// example: A123456789
	// required: true
	ID string `json:"id"`
	// The user healthCardID
	// example: 000011112222
	// required: true
	HealthCardID string `json:health_card_id`
	// The user birthday
	// example: 2022/05/23
	// required: true
	BirthDay string `json:"birthday"`
	// The user address
	// example: No.1, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan
	// required: true
	Address string `json:"address"`
	// The user phone number
	// example: 0912345678
	// required: true
	Phone string `json:"phone"`
	// The user vaccines
	Vaccines []string `json:vaccines`
}

func GetUser(client *vaccineclient.VaccineClient, nationID string) (*UserModel, error) {
	row, err := client.GetUser(nationID)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, nil
	}
	user := &UserModel{
		ID: nationID,
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
			// TODO: Handle unknown field error
		}
	}
	return user, nil
}
