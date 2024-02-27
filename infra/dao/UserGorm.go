package dao

import "gorm.io/gorm"

type UserGorm struct {
	gorm.Model
	Email          string `gorm:"unique"`
	HashedPassword string
}

func (UserGorm) TableName() string {
	return "users"
}
