package handlers

import (
	"gorm.io/gorm"
)

type AppConfig struct {
	Client *gorm.DB
}

func SetConfig(db *gorm.DB) *AppConfig {
	return &AppConfig{
		Client: db,
	}
}

var App *AppConfig
