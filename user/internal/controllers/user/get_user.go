package user

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/apierrors"
	"github.com/gin-gonic/gin"
)

// GetUserByID returns the user information.
// swagger:route GET /api/users/:nation_id User GetUser
//
// Get the user information.
//
// Responses:
//   200: UserResponse
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   404: NotFoundErrorResponse
//   500: InternalServerErrorResponse
//
func (u *User) GetUserByID(c *gin.Context) {
	nationID := c.Param("nation_id")
	// TODO: verify nationID in jwt token is the same as this nationID, otherwise return unauthorized

	row, err := u.vaccineClient.GetUser(nationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get user: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("user %s not found", nationID)))
		return
	}
	user, err := ConvertRowToUserModel(nationID, row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to user: %w", err)))
		return
	}

	c.JSON(http.StatusOK, user)
}
