package session

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/user"
	"github.com/gin-gonic/gin"
)

// GetSession returns the user information.
// swagger:route GET /api/session Session GetSession
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
func (s *Session) GetSession(c *gin.Context) {
	nationID, exists := c.Get("nationID")
	if !exists {
		_ = c.Error(apierrors.NewUnauthorizedError(errors.New("no token provided")))
		return
	}

	row, err := s.vaccineClient.GetUser(nationID.(string))
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get user: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("user %s not found", nationID)))
		return
	}
	user, err := user.ConvertRowToUserModel(nationID.(string), row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to user: %w", err)))
		return
	}

	c.JSON(http.StatusOK, user)
}
