package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/user"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// Options collects units to run this service.
type Options struct {
	Env                      env.Environments
	errorCollectorMiddleware gin.HandlerFunc
	userController           user.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	vaccineClient := vaccineclient.NewVaccineClient(env.ProjectID, env.InstanceID, env.TableName)

	userOpt := user.Options{
		Env:           env,
		VaccineClient: vaccineClient,
	}

	return &Options{
		Env:                      env,
		errorCollectorMiddleware: middlewares.NewErrorCollectorMiddleware(),
		userController:           user.New(userOpt),
	}, nil
}
