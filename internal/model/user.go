package model

import (
	"gorm.io/gorm"
)

type User struct {
	Id    uint
	Email string
	// Notes
	// Subscriptions
	// Collections
}

type UserGorm struct {
	gorm.Model
	Email          string `gorm:"unique"`
	HashedPassword string
}

func (UserGorm) TableName() string {
	return "users"
}
