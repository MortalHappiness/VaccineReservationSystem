package models

import (
	"fmt"
	"strings"

	"cloud.google.com/go/bigtable"
)

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
	// example: ["AZ", "BNT"]
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
		case "birthDay":
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

// ConvertUserModelToRow converts *UserModel to attributes of bigtable.Row.
func ConvertUserModelToAttributes(nationID string, user *UserModel) map[string]string {
	attributes := map[string]string{}
	if user.Name != "" {
		attributes["name"] = user.Name
	}
	if user.Gender != "" {
		attributes["gender"] = user.Gender
	}
	if user.HealthCardID != "" {
		attributes["healthCardID"] = user.HealthCardID
	}
	if user.BirthDay != "" {
		attributes["birthDay"] = user.BirthDay
	}
	if user.Address != "" {
		attributes["address"] = user.Address
	}
	if user.Phone != "" {
		attributes["phone"] = user.Phone
	}
	attributes["vaccines"] = strings.Join(user.Vaccines, ",")
	return attributes
}
