package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Environment string         `json:"environment"`
	Server      ServerConfig   `json:"server_config"`
	Mongo       MongoConfig    `json:"mongo_config"`
	RabbitMQ    RabbitMQConfig `json:"rabbitmq_config"`
	Redis       RedisConfig    `json:"redis_config"`
}

type ServerConfig struct {
	Port    int    `json:"port"`
	AppName string `json:"app_name"`
}

type MongoConfig struct {
	ConnectionURL          string `json:"connection_url"`
	DatabaseName           string `json:"database_name"`
	ServerSelectionTimeout int    `json:"server_selection_timeout"`
	MaxPoolSize            int    `json:"max_pool_size"`
	MinPoolSize            int    `json:"min_pool_size"`
	MaxConnectionIdleTime  int    `json:"max_connection_ideal_time"`
	ConnectTimeout         int    `json:"connect_timeout"`
	SocketTimeout          int    `json:"socket_timeout"`
}

type RabbitMQConfig struct {
	Protocol               string `json:"protocol"`
	BrokerHost             string `json:"broker_host"`
	BrokerPort             int    `json:"broker_port"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	WaitDurationForPublish string `json:"wait_duration_for_publish"`
	QoSLevel               int    `json:"qos_level"`
}

type RedisConfig struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	DB            int    `json:"db"`
	PoolSize      int    `json:"pool_size"`
	MinIdealConns int    `json:"min_ideal_conns"`
	UseTLS        bool   `json:"use_tls"`
}

// Load config/app.config.json
func LoadConfig() (*Config, error) {
	const configPath = "config/app.config.json"

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", configPath, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}
