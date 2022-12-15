package handlers

import (
	"gorm.io/gorm"
)

type AppConfig struct {
	client *gorm.DB
}

func SetConfig(db *gorm.DB) *AppConfig {
	return &AppConfig{
		client: db,
	}
}

var App *AppConfig
