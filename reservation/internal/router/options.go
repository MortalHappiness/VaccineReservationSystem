package router

import (
	"github.com/MortalHappiness/VaccineReservationSystem/reservation/internal/controllers/reservation"
	"github.com/MortalHappiness/VaccineReservationSystem/reservation/internal/env"
)

// Options collects units to run this service.
type Options struct {
	Env                   env.Environments
	reservationController reservation.I
}

// NewOptions returns Options based on environment variables.
//nolint:unparam // ignore unused parameter
func NewOptions(env env.Environments) (*Options, error) {
	reservationOpt := reservation.Options{
		Env: env,
	}

	return &Options{
		Env:                   env,
		reservationController: reservation.New(reservationOpt),
	}, nil
}
