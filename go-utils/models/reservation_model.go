package models

// ReservationModel is the body format of ReservationResponse
//
// swagger:model ReservationModel
type ReservationModel struct {
	// The reservation id
	// required: true
	// example: 0001
	ID string `json:"id" binding:"required"`
	// The reservation of the user
	// example: Bob
	// required: true
	User *UserModel `json:"user"`
	// The reservation hospital
	// example: Taipei City Hospital Heping Fuyou Branch
	// required: true
	Hospital *HospitalModel `json:"hospital"`
	// The reservation vaccinetype
	// example: BNT
	// required: true
	VaccineType string `json:"vaccinetype"`
	// The reservation date
	// example: 1653974953
	// required: true
	Date int64 `json:"date"`
	// The vaccination is completed
	// example: false
	// required: true
	Completed bool `json:"completed"`
}
