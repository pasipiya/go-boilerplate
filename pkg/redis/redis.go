package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"go-boilerplate/config"

	goredis "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type RedisClient struct {
	Client *goredis.Client
}

func NewRedisClient(cfg config.RedisConfig) (*RedisClient, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	options := &goredis.Options{
		Addr:         addr,
		Username:     cfg.Username,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdealConns,
	}

	if cfg.UseTLS {
		options.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	client := goredis.NewClient(options)

	// Test connection
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisClient{Client: client}, nil
}

func (r *RedisClient) Close() {
	_ = r.Client.Close()
}
