package models

import (
	"fmt"
	"strconv"
	"strings"

	"cloud.google.com/go/bigtable"
)

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
	// The reservation vaccineType
	// example: BNT
	// required: true
	VaccineType string `json:"vaccineType"`
	// The reservation date
	// example: 1653974953
	// required: true
	Date int64 `json:"date"`
	// The vaccination is completed
	// example: false
	// required: true
	Completed bool `json:"completed"`
}

func ConvertRowToReservationModel(rowKey string, row bigtable.Row) (*ReservationModel, error) {
	// parse rowKey
	rowKeyParts := strings.Split(rowKey, "#")
	if len(rowKeyParts) != 3 {
		return nil, fmt.Errorf("invalid rowKey %s", rowKey)
	}

	reservation := &ReservationModel{
		ID: rowKeyParts[2],
		User: &UserModel{
			NationID: rowKeyParts[1],
		},
		Hospital: &HospitalModel{},
	}

	for _, col := range row["reservation"] {
		qualifier := col.Column[strings.IndexByte(col.Column, ':')+1:]
		switch qualifier {
		case "county":
			reservation.Hospital.County = string(col.Value)
		case "township":
			reservation.Hospital.Township = string(col.Value)
		case "hospitalID":
			reservation.Hospital.ID = string(col.Value)
		case "vaccineType":
			reservation.VaccineType = string(col.Value)
		case "date":
			reservation.Date, _ = strconv.ParseInt(string(col.Value), 10, 64)
		case "completed":
			reservation.Completed = string(col.Value) == "true"
		}
	}

	return reservation, nil
}

func ConvertReservationModelToAttributes(reservation *ReservationModel) map[string]string {
	attributes := map[string]string{}
	if reservation.Hospital != nil {
		attributes["hospitalID"] = reservation.Hospital.ID
		if reservation.Hospital.County != "" {
			attributes["county"] = reservation.Hospital.County
		}
		if reservation.Hospital.Township != "" {
			attributes["township"] = reservation.Hospital.Township
		}
	}
	if reservation.VaccineType != "" {
		attributes["vaccineType"] = reservation.VaccineType
	}
	if reservation.Date != 0 {
		attributes["date"] = strconv.FormatInt(reservation.Date, 10)
	}
	attributes["completed"] = strconv.FormatBool(reservation.Completed)

	return attributes
}
