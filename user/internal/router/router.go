package router

import (
	"github.com/gin-gonic/gin"
)

// MakeRouter implements API SPEC.
func MakeRouter(opt *Options) *gin.Engine {
	var router *gin.Engine
	if opt.Env.AccessLog {
		router = gin.Default()
	} else {
		router = gin.New()
		router.Use(gin.Recovery())
	}

	api := router.Group("/api")
	{
		api.Use(opt.errorCollectorMiddleware)
		// users
		api.GET("/users/:nationID", opt.authMiddleware, opt.userController.GetUserByID)
		api.POST("/users", opt.userController.PostUser)
		api.PUT("/users/:nationID", opt.authMiddleware, opt.userController.PutUser)
		api.DELETE("/users/:nationID", opt.userController.DelUser)
		// session
		api.GET("/session", opt.authMiddleware, opt.sessionController.GetSession)
		api.POST("/session", opt.sessionController.PostSession)
		api.DELETE("/session", opt.authMiddleware, opt.sessionController.DelSession)
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
