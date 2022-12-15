package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Category    datatypes.JSON
	CreatedBy   uint
	Description string
	End         time.Time
	Start       time.Time
	Title       string
}

type User struct {
	gorm.Model
	Birthday time.Time
	Email    string
	Name     string
}