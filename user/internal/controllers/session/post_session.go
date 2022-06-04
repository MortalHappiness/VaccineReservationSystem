package session

import (
	"fmt"
	"net/http"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/jwt"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/user"
	"github.com/gin-gonic/gin"
)

// PostSession lets user login.
// swagger:route POST /api/sessions Session PostSessionRequest
//
// Login.
//
// Responses:
//   200: SessionResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (s *Session) PostSession(c *gin.Context) {
	var session PostSessionRequestModel
	err := c.ShouldBindJSON(&session)
	if err != nil {
		_ = c.Error(apierrors.NewBadRequestError(err))
		return
	}

	// get user information
	row, err := s.vaccineClient.GetUser(session.NationID)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to get user: %w", err)))
		return
	}
	if row == nil {
		_ = c.Error(apierrors.NewNotFoundError(fmt.Errorf("user %s not found", session.NationID)))
		return
	}
	user, err := user.ConvertRowToUserModel(session.NationID, row)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to convert row to user: %w", err)))
		return
	}

	// verify healthCardID
	if user.HealthCardID != session.HealthCardID {
		_ = c.Error(apierrors.NewUnauthorizedError(fmt.Errorf("healthCardID is not correct")))
		return
	}

	// set cookie
	token, err := jwt.NewToken(user.NationID, s.env.Secret)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(fmt.Errorf("failed to generate JWT token: %w", err)))
		return
	}
	c.SetCookie("token", token, 0, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}

// PostSessionRequest is the request of PostSession.
//
// swagger:parameters PostSessionRequest
type PostSessionRequest struct {
	// in: body
	// required: true
	Body *PostSessionRequestModel `json:"body"`
}

// PostSessionRequestModel is the body format of PostSessionRequest.
//
// swagger:model PostSessionRequestModel
type PostSessionRequestModel struct {
	// The user's nationID
	// example: A123456789
	// required: true
	NationID string `json:"nationID" binding:"required"`
	// The user's healthCardID
	// example: 000011112222
	// required: true
	HealthCardID string `json:"healthCardID" binding:"required"`
}
