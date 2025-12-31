package app

import (
	"github.com/gin-gonic/gin"

	healthController "go-boilerplate/internal/health/controller"
	healthRoutes "go-boilerplate/internal/health/routes"
)

type Router struct {
	HealthController *healthController.HealthController
}

func (r *Router) Register(app *gin.Engine) {
	api := app.Group("/api")

	// Health domain routes
	healthRoutes.RegisterHealthRoutes(api, r.HealthController)
}
