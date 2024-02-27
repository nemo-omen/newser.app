package repository

import (
	"gorm.io/gorm"
	"newser.app/model"
)

type NewsfeedGormRepo struct {
	DB *gorm.DB
}

func NewNewsfeedGormRepo(db *gorm.DB) NewsfeedGormRepo {
	return NewsfeedGormRepo{DB: db}
}

func (r NewsfeedGormRepo) Get(id uint) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r NewsfeedGormRepo) Create(n model.Newsfeed) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r NewsfeedGormRepo) All() []model.Newsfeed {
	return []model.Newsfeed{}
}

func (r NewsfeedGormRepo) Update(n model.Newsfeed) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r NewsfeedGormRepo) Delete(id uint) error {
	return nil
}

func (r NewsfeedGormRepo) FindBySlug(slug string) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}
