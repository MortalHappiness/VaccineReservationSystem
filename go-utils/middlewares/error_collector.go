package middlewares

import (
	"errors"

	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/apierrors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewErrorCollectorMiddleware returns a middleware which collects errors.
func NewErrorCollectorMiddleware() gin.HandlerFunc {
	log := logrus.New()
	handleError := func(c *gin.Context, err *apierrors.APIError) {
		// log error and response
		log.WithFields(logrus.Fields{
			"timestamp": err.Timestamp,
			"status":    err.Status,
			"error":     err.ErrMsg,
			"message":   err.Message,
		}).Error("error response")

		// write status code and content to response
		c.AbortWithStatusJSON(err.Status, err)
	}

	return func(c *gin.Context) {
		c.Next()
		// Collect errors after controller finish.
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				var apiErr *apierrors.APIError
				if errors.As(err.Err, &apiErr) {
					handleError(c, apiErr)
					break
				}
			}
		}
	}
}
