package worker

import (
	"fmt"

	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/env"
	"github.com/MortalHappiness/VaccineReservationSystem/user/internal/router"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

// Worker is the instance of this service.
type Worker struct {
	Router *gin.Engine
	Port   string
}

// New creates and setup the instance of this service.
func New() *Worker {
	e := env.Env
	opt, err := router.NewOptions(e)
	if err != nil {
		panic(fmt.Errorf("failed to create options: %w", err))
	}

	// print environment variables
	scs := spew.ConfigState{DisableCapacities: true, DisableMethods: true}
	scs.Dump(opt.Env)

	return &Worker{
		Router: router.MakeRouter(opt),
		Port:   e.Port,
	}
}

// Run is the method to launch HTTP server for this service.
func (a *Worker) Run() {
	_ = a.Router.Run(fmt.Sprintf(":%v", a.Port))
}
