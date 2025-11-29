// @title Maruoi AutoHub API
// @version 1.0
// @description API documentation for Maruoi AutoHub
// @host localhost:8080
// @BasePath /
package main

import (
	"log"

	_ "go-boilerplate/docs"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/routes"
	"go-boilerplate/pkg/db"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cfg := config.LoadConfig()

	mongoClient, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		log.Fatalf("❌ Mongo connection failed: %v", err)
	}

	routes.RegisterRoutes(router, mongoClient)

	router.Run(":8080")
}
