package middlewares

import (
	"errors"

	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/apierrors"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/jwt"
	"github.com/gin-gonic/gin"
)

// NewAuthMiddleware returns a middleware which authenticates user.
func NewAuthMiddleware(e env.Environments) gin.HandlerFunc {
	return func(c *gin.Context) {
		// read token from cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			_ = c.Error(apierrors.NewUnauthorizedError(errors.New("no token provided")))
			c.Abort()
			return
		}

		// Parse token.
		token, err := jwt.ParseToken(tokenString, e.Secret)
		if err != nil {
			_ = c.Error(apierrors.NewUnauthorizedError(errors.New("token is invalid")))
			c.Abort()
			return
		}
		// Set nationID to context.
		c.Set("nationID", token.NationID)
		c.Next()
	}
}
