package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/controllers/infoexample"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
)

// Options collects units to run this service.
type Options struct {
	Env            env.Environments
	infoController infoexample.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	infoOpt := infoexample.Options{
		Env: env,
	}

	return &Options{
		Env:            env,
		infoController: infoexample.New(infoOpt),
	}, nil
}
