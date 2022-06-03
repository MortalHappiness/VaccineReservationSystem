package user

import (
	"fmt"
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

	// nationID and healthCardID is required in POST request
	if user.NationID == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("nationID is empty")))
		return
	}
	if user.HealthCardID == "" {
		_ = c.Error(apierrors.NewBadRequestError(fmt.Errorf("healthCardID is empty")))
		return
	}

	// TODO: user should return user
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

	// TODO: user should be changed to the return value of CreateOrUpdateUser
	c.JSON(http.StatusOK, user)
}

// PostUserRequest is the request of PostUser
//
// swagger:parameters PostUserRequest
type PostUserRequest struct {
	// The user info
	// in: body
	User *UserModel `json:"user"`
}
