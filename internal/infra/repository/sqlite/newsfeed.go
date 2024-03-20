package sqlite

import (
	"github.com/jmoiron/sqlx"
	"newser.app/internal/dto"
)

type NewsfeedSqliteRepo struct {
	db *sqlx.DB
}

func NewNewsfeedSqliteRepo(db *sqlx.DB) *NewsfeedSqliteRepo {
	return &NewsfeedSqliteRepo{
		db: db,
	}
}

func (r *NewsfeedSqliteRepo) GetNewsfeed(email string) (*dto.NewsfeedDTO, error) {
	// nf := &dto.NewsfeedDTO{}
	return nil, nil
}
