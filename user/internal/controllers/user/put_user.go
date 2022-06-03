package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/apierrors"
	"github.com/gin-gonic/gin"
)

// PutUser updates a user and returns him/her.
// swagger:route PUT /api/users/:nationID User PutUserRequest
//
// Update a user.
//
// Responses:
//   200: UserResponse
//   400: BadRequestErrorResponse
//   401: UnauthorizedErrorResponse
//   500: InternalServerErrorResponse
//
func (u *User) PutUser(c *gin.Context) {
	nationID := c.Param("nationID")

	var user UserModel
	err := c.ShouldBindJSON(&user)
	if err != nil {
		_ = c.Error(apierrors.NewBadRequestError(err))
		return
	}

	// verify user auth
	err = u.AuthVerify(c, nationID)
	if err != nil {
		_ = c.Error(apierrors.NewUnauthorizedError(err))
		return
	}

	// verify nationID is the same as the one in the request
	if nationID != user.NationID {
		_ = c.Error(apierrors.NewBadRequestError(
			fmt.Errorf("nationID is not matched: %s != %s", nationID, user.NationID)))
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

// PutUserRequest is the request of PutUser
//
// swagger:parameters PutUserRequest
type PutUserRequest struct {
	// The user info
	// in: body
	User *UserModel `json:"user"`
}
