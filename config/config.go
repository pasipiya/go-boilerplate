package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds the configuration variables
type Config struct {
	ServerPort string
	PprofPort  string
}

// LoadConfig loads environment variables from the .env file and returns a Config struct
func LoadConfig() *Config {
	// Load .env file, ignore error if it’s missing but log if there’s an actual error loading
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v\n", err)
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		PprofPort:  getEnv("PPROF_PORT", "6060"),
	}
}

// getEnv retrieves the value of the environment variable `key` or returns `fallback` if not set
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
