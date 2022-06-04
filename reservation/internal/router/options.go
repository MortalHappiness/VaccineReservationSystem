package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/middlewares"
	"github.com/MortalHappiness/VaccineReservationSystem/reservation/internal/controllers/reservation"
	"github.com/MortalHappiness/VaccineReservationSystem/reservation/internal/env"
	"github.com/gin-gonic/gin"
)

// Options collects units to run this service.
type Options struct {
	Env                      env.Environments
	errorCollectorMiddleware gin.HandlerFunc
	authMiddleware           gin.HandlerFunc
	reservationController    reservation.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	vaccineClient := vaccineclient.NewVaccineClient(env.ProjectID, env.InstanceID, env.TableName)

	reservationOpt := reservation.Options{
		Env:           env,
		VaccineClient: vaccineClient,
	}

	return &Options{
		Env:                      env,
		errorCollectorMiddleware: middlewares.NewErrorCollectorMiddleware(),
		authMiddleware:           middlewares.NewAuthMiddleware(env.Secret),
		reservationController:    reservation.New(reservationOpt),
	}, nil
}
