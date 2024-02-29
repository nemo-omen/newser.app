package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type NewsfeedSqliteRepo struct {
	DB *sqlx.DB
}

func NewNewsfeedSqliteRepo(db *sqlx.DB) NewsfeedSqliteRepo {
	return NewsfeedSqliteRepo{DB: db}
}

func (r NewsfeedSqliteRepo) Migrate() error {
	fmt.Println("Starting newsfeed table migration...")
	qa := `
	CREATE TABLE IF NOT EXISTS newsfeeds(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		site_url TEXT NOT NULL,
		feed_url TEXT NOT NULL,
		description TEXT,
		image JSON,
		published TEXT NOT NULL,
		published_parsed DATETIME NOT NULL,
		updated TEXT NOT NULL,
		updated_parsed DATETIME NOT NULL,
		copyright TEXT,
		author_id INT,
		language TEXT,
		feed_type TEXT,
		slug TEXT NOT NULL
	);
	`
	_, err := r.DB.Exec(qa)
	if err != nil {
		fmt.Println("error migrating newsfeeds: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating newsfeeds")
	}
	return err
}

func (r NewsfeedSqliteRepo) Get(id uint) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r NewsfeedSqliteRepo) Create(n model.Newsfeed) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r NewsfeedSqliteRepo) All() []model.Newsfeed {
	return []model.Newsfeed{}
}

func (r NewsfeedSqliteRepo) Update(n model.Newsfeed) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}

func (r NewsfeedSqliteRepo) Delete(id uint) error {
	return nil
}

func (r NewsfeedSqliteRepo) FindBySlug(slug string) (model.Newsfeed, error) {
	return model.Newsfeed{}, nil
}
