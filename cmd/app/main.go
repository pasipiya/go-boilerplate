package main

import (
	"log"

	"github.com/pasipiya/go-boilerplate/config"
	"github.com/pasipiya/go-boilerplate/internal/server"
	"github.com/pasipiya/go-boilerplate/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.Info("This is a info message")

	// Start HTTP server
	if err := server.Start(cfg); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
