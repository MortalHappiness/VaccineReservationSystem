package infoexample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInfoV1 returns Info.
// swagger:route Get /v1/info InfoExample GetInfo
//
// Get Information
//
// GetInfo returns some information.
//
// Responses:
//   200: infoResponse
//   500: InternalServerErrorResponse
//
func (a *InfoExample) GetInfoV1(c *gin.Context) {
	c.JSON(http.StatusOK, InfoModel{Name: "Trump"})
}

// InfoResponse ...
//
// swagger:response infoResponse
type InfoResponse struct {
	// The user info
	// in: body
	Info *InfoModel `json:"info"`
}
