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
	router.Use(CORSMiddleware())
	// exampleMiddleware required group
	api := router.Group("/api")
	{
		api.Use(opt.errorCollectorMiddleware)
		api.Use(opt.authMiddleware)
		api.GET("/reservations/users/:nationID", opt.reservationController.GetReservation)
		api.POST("/reservations/users/:nationID", opt.reservationController.PostReservation)
		api.DELETE("/reservations/users/:nationID/:reservationID", opt.reservationController.DeleteReservation)
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
			//   404: genericErrorcd
			//
			router.Static("/.spec", opt.Env.SpecFiles)
		}
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
