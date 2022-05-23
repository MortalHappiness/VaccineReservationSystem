package user

import (
	"net/http"

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
	model := UserModel{
		Name:     "bob",
		Gender:   "male",
		ID:       "A123456789",
		BirthDay: "2022/05/23",
		Address:  "No.1, Sec. 4, Roosevelt Road, Taipei, 10617 Taiwan",
		Phone:    "0912345678",
	}
	c.JSON(http.StatusOK, model)
}
