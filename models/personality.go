package models

import (
	"time"

	"gorm.io/gorm"
)

type Personality struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	Name      string
	History   string
	UpdatedAt time.Time
	CreatedAt time.Time
}

var Personalities []Personality
