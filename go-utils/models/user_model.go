package models

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
