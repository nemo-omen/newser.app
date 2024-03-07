package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
	"newser.app/shared/util"
)

type NewsfeedSqliteRepo struct {
	db *sqlx.DB
}

func NewNewsfeedSqliteRepo(db *sqlx.DB) *NewsfeedSqliteRepo {
	return &NewsfeedSqliteRepo{db: db}
}

func (r *NewsfeedSqliteRepo) Migrate() error {
	fmt.Println("Starting newsfeed table migration...")
	qa := `
	CREATE TABLE IF NOT EXISTS newsfeeds(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		site_url NOT NULL,
		feed_url TEXT UNIQUE NOT NULL,
		description TEXT,
		copyright TEXT,
		language TEXT,
		feed_type TEXT,
		slug TEXT NOT NULL
	);
	`
	_, err := r.db.Exec(qa)
	if err != nil {
		fmt.Println("error migrating newsfeeds: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating newsfeeds")
	}
	return err
}

func (r *NewsfeedSqliteRepo) Get(id int64) (*model.Newsfeed, error) {
	nf := model.Newsfeed{}
	err := r.db.Get(&nf, `SELECT * FROM newsfeeds WHERE id=?`, id)
	if err != nil {
		fmt.Println("error getting feed: ", err.Error())
		return nil, err
	}
	return &nf, nil
}

func (r *NewsfeedSqliteRepo) Create(n *model.Newsfeed) (*model.Newsfeed, error) {
	q := `
	INSERT INTO newsfeeds(
		title,
		site_url,
		feed_url,
		description,
		copyright,
		language,
		feed_type,
		slug
	)
		VALUES(?,?,?,?,?,?,?,?)
		ON CONFLICT(feed_url) do nothing;
	`
	res, err := r.db.Exec(
		q,
		n.Title,
		n.SiteUrl,
		n.FeedUrl,
		n.Description,
		n.Copyright,
		n.Language,
		n.FeedType,
		util.Slugify(n.Title),
	)

	if err != nil {
		fmt.Println("article insert error: ", err.Error())
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	n.ID = id
	return n, nil
}

func (r *NewsfeedSqliteRepo) All() []*model.Newsfeed {
	return nil
}

func (r *NewsfeedSqliteRepo) Update(n *model.Newsfeed) (*model.Newsfeed, error) {
	return nil, nil
}

func (r *NewsfeedSqliteRepo) Delete(id int64) error {
	return nil
}

func (r *NewsfeedSqliteRepo) FindBySlug(slug string) (*model.Newsfeed, error) {
	return nil, nil
}
