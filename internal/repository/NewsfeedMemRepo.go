package repository

import "newser.app/internal/model"

type NewsfeedMemRepo struct {
	Feeds []model.Newsfeed
}

func (r *NewsfeedMemRepo) Get(id uint) model.Newsfeed {
	return model.Newsfeed{}
}

func (r *NewsfeedMemRepo) Create(n model.Newsfeed) (uint, error) {
	return 0, nil
}

func (r *NewsfeedMemRepo) All() []model.Newsfeed {
	return []model.Newsfeed{}
}

func (r *NewsfeedMemRepo) Update(m model.Newsfeed) (model.Newsfeed, error) {
	return m, nil
}

func (r *NewsfeedMemRepo) Delete(id uint) error {
	return nil
}
