package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	PoolMin  uint64
	PoolMax  uint64
}

type Config struct {
	Mongo MongoConfig
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}

	poolMin := parseUint(getEnv("MONGO_POOL_MIN", "5"))
	poolMax := parseUint(getEnv("MONGO_POOL_MAX", "50"))

	return &Config{
		Mongo: MongoConfig{
			Host:     getEnv("MONGO_HOST", "localhost"),
			Port:     getEnv("MONGO_PORT", "27017"),
			User:     getEnv("MONGO_USER", ""),
			Pass:     getEnv("MONGO_PASS", ""),
			Database: getEnv("MONGO_DATABASE", ""),
			PoolMin:  poolMin,
			PoolMax:  poolMax,
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func parseUint(val string) uint64 {
	v, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return v
}
