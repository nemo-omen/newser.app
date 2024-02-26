package repository

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"newser.app/internal/model"
)

type UserGormRepo struct {
	DB *gorm.DB
}

func NewUserGormRepo(dsn string) UserGormRepo {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	return UserGormRepo{DB: db}
}

func (r UserGormRepo) Get(id uint) (model.User, error) {
	return model.User{}, nil
}

func (r UserGormRepo) Create(ug model.UserGorm) (model.User, error) {
	// res := r.DB.Create(&ug)
	return model.User{}, nil
}

func (r UserGormRepo) All() []model.User {
	return []model.User{}
}

func (r UserGormRepo) Update(m model.UserGorm) (model.User, error) {
	return model.User{}, nil
}

func (r UserGormRepo) Delete(id uint) error {
	return nil
}
