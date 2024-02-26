package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint
	Email     string
	CreatedAt time.Time
	// Notes
	// Subscriptions
	// Collections
}

type UserGorm struct {
	gorm.Model
	Email          string `gorm:"unique"`
	HashedPassword string
}
