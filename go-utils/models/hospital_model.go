package models

// Hospital model is the body format of HospitalResponse
//
// swagger:model HospitalModel
type HospitalModel struct {
	// The hospital id
	// required: true
	// example: 0001
	ID string `json:"id" binding:"required"`
	// The hospital name
	// example: Taipei City Hospital Heping Fuyou Branch
	// required: true
	Name string `json:"name"`
	// The hospital County
	// example: New Taipei City
	// required: true
	County string `json:"county"`
	// The hospital Township
	// example: Banqiao
	// required: true
	Township string `json:"township"`
	// The hospital address
	// example: No.1, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan
	// required: true
	Address string `json:"address"`
	// The hospital vaccines
	// example: { "AZ": 100, "BNT": 200 }
	// required: true
	VaccineCnt map[string]int `json:"vaccineCnt"`
}
