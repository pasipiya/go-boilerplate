package db

import (
	"context"
	"fmt"
	"time"

	"go-boilerplate/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg config.MongoConfig) (*mongo.Client, error) {
	// Build URI: supports BOTH with and without authentication
	var uri string

	if cfg.User != "" && cfg.Pass != "" {
		// Authenticated
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s?authSource=admin",
			cfg.User,
			cfg.Pass,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
	} else {
		// No authentication
		uri = fmt.Sprintf(
			"mongodb://%s:%s/%s",
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
	}

	clientOpts := options.Client().
		ApplyURI(uri).
		SetMinPoolSize(cfg.PoolMin).
		SetMaxPoolSize(cfg.PoolMax)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
