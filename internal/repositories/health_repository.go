package repositories

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

type HealthRepository struct {
    mongo *mongo.Client
}

func NewHealthRepository(mongoClient *mongo.Client) *HealthRepository {
    return &HealthRepository{mongo: mongoClient}
}

func (r *HealthRepository) PingDB() error {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    return r.mongo.Ping(ctx, readpref.Primary())
}
