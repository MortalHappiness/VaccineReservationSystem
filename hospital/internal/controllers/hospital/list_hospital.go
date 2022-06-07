package hospital

// GetHospital returns the hospital information.
// swagger:route GET /api/hospitals Hospital GetHospital
//
// Get the hospital information.
// Parameters:
//   + name: county
//     in: query
//     description: the county of the hospital
//     required: true
//     type: string
//   + name: township
//     in: query
//     description: the township of the hospital
//     required: true
//     type: string
//
// Responses:
//   200: HospitalResponse
//	 400: BadRequestErrorResponse
//   404: NotFoundErrorResponse
//   500: InternalServerErrorResponse
//
// func (u *Hospital) ListHospitals(c *gin.Context) {
// 	// get hospitals
// 	rows, err := u.vaccineClient.ListHospitals()
// 	if err != nil {
// 		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get hospitals: %w", err)))
// 		return
// 	}
// 	if len(rows) == 0 {
// 		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("no hospital found")))
// 		return
// 	}
// 	hospitals := []models.HospitalModel{}
// 	for _, row := range rows {
// 		// parse row key ID
// 		hospital := &models.HospitalModel{}
// 		hospital, err = models.ConvertRowToHospitalModel(row.Key(), row)
// 		if err != nil {
// 			_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to hospital: %w", err)))
// 			return
// 		}
// 		hospitals = append(hospitals, *hospital)
// 	}

// 	c.JSON(http.StatusOK, hospitals)
// }
