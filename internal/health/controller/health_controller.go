package controllers

import (
	"net/http"

	services "go-boilerplate/internal/health/service"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	service *services.HealthService
}

func NewHealthController(service *services.HealthService) *HealthController {
	return &HealthController{service: service}
}

// Health godoc
// @Summary Health check endpoint
// @Description API and DB health info
// @Tags Health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /api/health [get]
func (h *HealthController) HealthCheck(c *gin.Context) {
	res := h.service.CheckHealth()
	c.JSON(http.StatusOK, res)
}
