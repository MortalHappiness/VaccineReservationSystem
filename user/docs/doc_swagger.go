// Package doc Worker.
//
// Documentation of Worker.
//
//     Schemes: https
//     Version: %VERSION_TO_BE_REPLACED_BY_GITRUNNER%
//     Host: localhost:7712
//     BasePath: /
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package doc

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// The error message
	//
	// required: true
	Message string `json:"message"`
}

// EmptyResponse are sent when the HTTP status code is 204.
//
// swagger:response emptyResponse
type EmptyResponse struct{}
