package bootstrap

import (
	controller "go-boilerplate/internal/health/controller"
	repository "go-boilerplate/internal/health/repository"
	service "go-boilerplate/internal/health/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	HealthController *controller.HealthController
}

func NewContainer(mongo *mongo.Client) *Container {
	// Health domain wiring
	healthRepo := repository.NewHealthRepository(mongo)
	healthService := service.NewHealthService(healthRepo)
	healthController := controller.NewHealthController(healthService)

	return &Container{
		HealthController: healthController,
	}
}
