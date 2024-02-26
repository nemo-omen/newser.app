package repository

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"newser.app/internal/model"
)

type NewsfeedGormRepo struct {
	DB *gorm.DB
}

func NewNewsfeedGormRepo(dsn string) NewsfeedGormRepo {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	return NewsfeedGormRepo{DB: db}
}

func (r *NewsfeedGormRepo) Get(id uint) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r *NewsfeedGormRepo) Create(ng model.NewsfeedGorm) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r *NewsfeedGormRepo) All() []model.Newsfeed {
	return []model.Newsfeed{}
}

func (r *NewsfeedGormRepo) Update(ng model.NewsfeedGorm) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r *NewsfeedGormRepo) Delete(id uint) error {
	return nil
}
