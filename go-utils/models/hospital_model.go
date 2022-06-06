package models

import (
	"encoding/json"
	"strings"

	"cloud.google.com/go/bigtable"
)

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

func ConvertRowToHospitalModel(ID string, row bigtable.Row) (*HospitalModel, error) {
	hosp := &HospitalModel{
		ID: ID,
	}
	for _, col := range row["hospital"] {
		qualifier := col.Column[strings.IndexByte(col.Column, ':')+1:]
		switch qualifier {
		case "name":
			hosp.Name = string(col.Value)
		case "county":
			hosp.County = string(col.Value)
		case "township":
			hosp.Township = string(col.Value)
		case "address":
			hosp.Address = string(col.Value)
		case "vaccineCnt":
			err := json.Unmarshal(col.Value, &hosp.VaccineCnt)
			if err != nil {
				return nil, err
			}
		}
	}
	return hosp, nil
}

func ConvertHospitalModelToAttributes(hospital *HospitalModel) map[string]string {
	vaccineCnt, _ := json.Marshal(hospital.VaccineCnt)
	attributes := map[string]string{}
	if hospital.Name != "" {
		attributes["name"] = hospital.Name
	}
	if hospital.County != "" {
		attributes["county"] = hospital.County
	}
	if hospital.Township != "" {
		attributes["township"] = hospital.Township
	}
	if hospital.Address != "" {
		attributes["address"] = hospital.Address
	}
	if hospital.VaccineCnt != nil {
		attributes["vaccineCnt"] = string(vaccineCnt)
	}
	return attributes
}
