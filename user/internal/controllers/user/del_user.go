package user

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/gin-gonic/gin"
)

// DelUser deletes a user and returns his/hers nation ID.
// swagger:route DELETE /api/users/:nationID User DelUserRequest
//
// Delete a user with his/her nation ID.
//
// Responses:
//   200: UserResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *User) DelUser(c *gin.Context) {
	nationID := c.Param("nationID")
	if nationID == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("nationID is empty")))
		return
	}

	err := u.vaccineClient.DeleteUser(nationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to delete user %s: %w", nationID, err)))
		return
	}

	c.JSON(http.StatusOK, nationID)
}

// DelUserRequest is the request of DelUser
//
// swagger:parameters DelUserRequest
type DelUserRequest struct {
	// The user info
	// in: body
	// example: "A123456789"
	NationID string `json:"nationID"`
}
