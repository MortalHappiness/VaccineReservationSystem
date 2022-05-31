package router

import (
	"github.com/gin-gonic/gin"
)

// MakeRouter implements API SPEC.
func MakeRouter(opt *Options) *gin.Engine {
	gin.DisableConsoleColor()
	var router *gin.Engine
	if opt.Env.AccessLog {
		router = gin.Default()
	} else {
		router = gin.New()
		router.Use(gin.Recovery())
	}
	// exampleMiddleware required group
	{
		// TODO: add apis
		v1 := router.Group("/v1")
		v1.GET("/hospital", opt.hospitalController.GetHospitalV1)
		v1.POST("/hospital", opt.hospitalController.PostHospitalV1)
		v1.PUT("/hospital/:id", opt.hospitalController.PutHospitalV1)
		v1.DELETE("/hospital/:id", opt.hospitalController.DeleteHospitalV1)
	}
	// API Spec Swagger UI
	{
		if opt.Env.SpecEnabled {
			// GetSpec returns API Spec. This is for testing & development usage.
			// swagger:route Get /.spec Development GetSpec
			//
			// Get Server API Specification
			//
			// GetSpec returns API Spec. This API is for testing & development only.
			//
			// Responses:
			//   301: emptyResponse
			//   404: genericError
			//
			router.Static("/.spec", opt.Env.SpecFiles)
		}
	}

	return router
}
