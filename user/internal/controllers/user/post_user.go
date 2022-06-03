package user

import (
	"net/http"
	"strings"

	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/apierrors"
	"github.com/gin-gonic/gin"
)

// PostUser adds a new user and returns him/her.
// swagger:route POST /api/users User PostUserRequest
//
// Add a new user.
//
// Responses:
//   200: UserResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *User) PostUser(c *gin.Context) {
	var user UserModel
	err := c.ShouldBindJSON(&user)
	if err != nil {
		_ = c.Error(apierrors.NewBadRequestError(err))
		return
	}
	err = u.vaccineClient.CreateOrUpdateUser(
		user.NationID,
		user.Name,
		user.HealthCardID,
		user.Gender,
		user.BirthDay,
		user.Address,
		user.Phone,
		strings.Join(user.Vaccines, ","),
	)
	if err != nil {
		_ = c.Error(apierrors.NewInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, user)
}

// PostUserRequest is the request of PostUserV1
//
// swagger:parameters PostUserRequest
type PostUserRequest struct {
	// The user info
	// in: body
	User *UserModel `json:"user"`
}
