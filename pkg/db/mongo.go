package db

import (
	"context"
	"time"

	"go-boilerplate/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg config.MongoConfig) (*mongo.Client, error) {
	clientOpts := options.Client().
		ApplyURI(cfg.ConnectionURL).
		SetMaxPoolSize(uint64(cfg.MaxPoolSize)).
		SetMinPoolSize(uint64(cfg.MinPoolSize)).
		SetServerSelectionTimeout(time.Duration(cfg.ServerSelectionTimeout) * time.Second).
		SetMaxConnIdleTime(time.Duration(cfg.MaxConnectionIdleTime) * time.Second).
		SetConnectTimeout(time.Duration(cfg.ConnectTimeout) * time.Second).
		SetSocketTimeout(time.Duration(cfg.SocketTimeout) * time.Second)

	// Global context timeout — 10s is good for initial connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	// Verify connectivity
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
