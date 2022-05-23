package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostUserV1 adds a new user and returns him/her.
// swagger:route POST /v1/user User PostUserRequest
//
// Add a new user.
//
// Responses:
//   200: UserResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (u *User) PostUserV1(c *gin.Context) {
	var model UserModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		_ = c.Error(fmt.Errorf("post user request: %w", err))
		return
	}
	c.JSON(http.StatusOK, model)
}

// PostUserRequest is the request of PostUserV1
//
// swagger:parameters PostUserRequest
type PostUserRequest struct {
	// The user info
	// in: body
	User *UserModel `json:"user"`
}
