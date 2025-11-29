package routes

import (
	"go-boilerplate/internal/controllers"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/services"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, mongoClient *mongo.Client) {

	// ----------------------------
	// HEALTH ENDPOINT WIRING
	// ----------------------------
	healthRepo := repositories.NewHealthRepository(mongoClient)
	healthService := services.NewHealthService(healthRepo)
	healthController := controllers.NewHealthController(healthService)
	HealthRoutes(router, healthController)

	// ----------------------------
	// You can attach more route groups here
	// ----------------------------
	// userRepo := repositories.NewUserRepository(mongoClient)
	// userService := services.NewUserService(userRepo)
	// userController := controllers.NewUserController(userService)
	// UserRoutes(router, userController)
}
