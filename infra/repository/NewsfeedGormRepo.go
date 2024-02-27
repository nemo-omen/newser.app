package repository

import (
	"gorm.io/gorm"
	"newser.app/infra/dao"
	"newser.app/model"
)

type NewsfeedGormRepo struct {
	DB *gorm.DB
}

func NewNewsfeedGormRepo(db *gorm.DB) NewsfeedGormRepo {
	return NewsfeedGormRepo{DB: db}
}

func (r *NewsfeedGormRepo) Get(id uint) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r *NewsfeedGormRepo) Create(ng dao.NewsfeedGorm) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r *NewsfeedGormRepo) All() []model.Newsfeed {
	return []model.Newsfeed{}
}

func (r *NewsfeedGormRepo) Update(ng dao.NewsfeedGorm) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r *NewsfeedGormRepo) Delete(id uint) error {
	return nil
}
