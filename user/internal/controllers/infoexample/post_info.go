package infoexample

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostInfoV1 returns Info.
// swagger:route Post /v1/info InfoExample PostInfo
//
// Post Information
//
// PostInfo returns some information.
// Responses:
//   204: emptyResponse
//   400: BadRequestErrorResponse
//   500: InternalServerErrorResponse
//
func (a *InfoExample) PostInfoV1(c *gin.Context) {
	var info InfoModel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		_ = c.Error(fmt.Errorf("post inforequest: %w", err))
		return
	}
	c.Status(http.StatusNoContent)
}

// PostInfoBody are body for POST info api.
//
// swagger:parameters PostInfo
type PostInfoBody struct {
	// The request body
	//
	// in: body
	// required: true
	Info *InfoModel `json:"info body"`
}
