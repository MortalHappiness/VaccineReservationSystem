package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DelSession deletes the session.
// swagger:route DELETE /api/session Session DelSession
//
// Logout the user.
//
// Responses:
//   204: NoContentResponse
//
func (s *Session) DelSession(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Status(http.StatusNoContent)
}
