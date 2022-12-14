package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"context"
)

type AppConfig struct {
	DB  *mongo.Client
	ctx context.Context
}

func SetConfig(db *mongo.Client, ctx context.Context) *AppConfig {
	return &AppConfig{
		DB:  db,
		ctx: ctx,
	}
}

var App *AppConfig
