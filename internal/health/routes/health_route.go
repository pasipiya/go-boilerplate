package routes

import (
	controller "go-boilerplate/internal/health/controller"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.RouterGroup, controller *controller.HealthController) {
	r.GET("/health", controller.HealthCheck)
}
