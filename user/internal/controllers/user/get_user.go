package user

import (
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/gin-gonic/gin"
)

// GetUserV1 returns the user information.
// swagger:route GET /v1/user User GetUser
//
// Get the user information.
//
// Responses:
//   200: UserResponse
//   500: InternalServerErrorResponse
//
func (u *User) GetUserV1(c *gin.Context) {
	// TODO: Read environment variables instead of hard code projectID and instanceID
	projectID := "my-project"
	instanceID := "my-instance"
	tableName := "vaccine-reservation-system"
	// TODO: Store vaccineClient somewhere to share between functions
	vaccineClient := vaccineclient.NewVaccineClient(projectID, instanceID, tableName)

	model, err := GetUser(vaccineClient, "A123456789")
	if err != nil {
		// TODO: Do error handling
	}
	c.JSON(http.StatusOK, model)
}
