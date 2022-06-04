package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/bigtable/pkg/vaccineclient"
	"github.com/MortalHappiness/VaccineReservationSystem/go-utils/middlewares"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/session"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/user"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/gin-gonic/gin"
)

// Options collects units to run this service.
type Options struct {
	Env                      env.Environments
	errorCollectorMiddleware gin.HandlerFunc
	authMiddleware           gin.HandlerFunc
	userController           user.I
	sessionController        session.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	vaccineClient := vaccineclient.NewVaccineClient(env.ProjectID, env.InstanceID, env.TableName)

	userOpt := user.Options{
		Env:           env,
		VaccineClient: vaccineClient,
	}

	sessionOpt := session.Options{
		Env:           env,
		VaccineClient: vaccineClient,
	}

	return &Options{
		Env:                      env,
		errorCollectorMiddleware: middlewares.NewErrorCollectorMiddleware(),
		authMiddleware:           middlewares.NewAuthMiddleware(env.Secret),
		userController:           user.New(userOpt),
		sessionController:        session.New(sessionOpt),
	}, nil
}
