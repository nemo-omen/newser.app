package repository

import (
	"database/sql"

	"newser.app/internal/model"
)

type NewsfeedSqliteRepo struct {
	DB *sql.DB
}

func (r *NewsfeedSqliteRepo) Get(id uint) model.Newsfeed {
	return model.Newsfeed{}
}

func (r *NewsfeedSqliteRepo) Create(n model.Newsfeed) (uint, error) {
	return 0, nil
}

func (r *NewsfeedSqliteRepo) All() []model.Newsfeed {
	return []model.Newsfeed{}
}

func (r *NewsfeedSqliteRepo) Update(m model.Newsfeed) (model.Newsfeed, error) {
	return m, nil
}

func (r *NewsfeedSqliteRepo) Delete(id uint) error {
	return nil
}
