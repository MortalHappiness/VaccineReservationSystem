package user

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/models"
	"github.com/gin-gonic/gin"
)

// GetUserByID returns the user information.
// swagger:route GET /api/users/:nationID User GetUser
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
	nationID := c.Param("nationID")

	// verify user auth
	err := AuthVerify(c, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
		return
	}

	// get user info
	row, err := u.vaccineClient.GetUser(nationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get user: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("user#%s not found", nationID)))
		return
	}
	user, err := models.ConvertRowToUserModel(nationID, row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to user: %w", err)))
		return
	}

	c.JSON(http.StatusOK, user)
}
