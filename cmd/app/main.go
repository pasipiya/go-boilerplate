// @title Go Boilerplate API
// @version 1.0
// @description API documentation for Go Boilerplate
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/internal/app"
	"go-boilerplate/internal/app/bootstrap"
	"go-boilerplate/pkg/db"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/rabbitmq"
	"go-boilerplate/pkg/redis"
	"strconv"

	_ "go-boilerplate/cmd/app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Load App Config
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to load config: %v", err))
		panic(err)
	}
	logger.Info("Configuration loaded successfully")

	// Initialize MongoDB
	mongoClient, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		logger.Error("MongoDB connection failed: " + err.Error())
		panic(err)
	}
	logger.Info("MongoDB connected successfully")

	// Initialize Redis
	redisClient, err := redis.NewRedisClient(cfg.Redis)
	if err != nil {
		logger.Error("Redis connection failed: " + err.Error())
		panic(err)
	}
	logger.Info("Redis connected successfully")
	defer redisClient.Close()

	// Initialize RabbitMQ
	rmq, err := rabbitmq.NewRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		logger.Error("RabbitMQ connection failed: " + err.Error())
		panic(err)
	}
	logger.Info("RabbitMQ connected successfully")
	defer rmq.Close()

	// Initialize Gin Router
	router := gin.Default()

	// Swagger Documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Dependency Injection Container
	container := bootstrap.NewContainer(mongoClient)
	logger.Info("Dependency container initialized")

	// Register Routes
	appRouter := app.Router{
		HealthController: container.HealthController,
	}
	appRouter.Register(router)
	logger.Info("Routes registered successfully")

	// Start Server
	serverAddr := ":" + strconv.Itoa(cfg.Server.Port)
	logger.Info(fmt.Sprintf(
		"Starting %s HTTP server at http://localhost%s",
		cfg.Server.AppName,
		serverAddr,
	))

	if err := router.Run(serverAddr); err != nil {
		logger.Error("Server failed: " + err.Error())
		panic(err)
	}
}
