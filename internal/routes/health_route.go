package routes

import (
	"go-boilerplate/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(router *gin.Engine, ctrl *controllers.HealthController) {
	router.GET("/health", ctrl.Health)
}
