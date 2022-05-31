package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/hospital/internal/controllers/hospital"
	"github.com/MortalHappiness/VaccineReservationSystem/hospital/internal/env"
)

// Options collects units to run this service.
type Options struct {
	Env                env.Environments
	hospitalController hospital.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	hospitalOpt := hospital.Options{
		Env: env,
	}

	return &Options{
		Env:                env,
		hospitalController: hospital.New(hospitalOpt),
	}, nil
}
