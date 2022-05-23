package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/user"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
)

// Options collects units to run this service.
type Options struct {
	Env            env.Environments
	userController user.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	userOpt := user.Options{
		Env: env,
	}

	return &Options{
		Env:            env,
		userController: user.New(userOpt),
	}, nil
}
