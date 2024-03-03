package repository

import (
	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type ImageRepository interface {
	Create(img *model.Image) (*model.Image, error)
	Get(id int64) (*model.Image, error)
	Migrate() error
}

type ImageSqliteRepository struct {
	db *sqlx.DB
}

func NewImageSqliteRepo(db *sqlx.DB) *ImageSqliteRepository {
	return &ImageSqliteRepository{
		db: db,
	}
}

func (r *ImageSqliteRepository) Create(i *model.Image) (*model.Image, error) {
	q := `
	INSERT INTO images(url, title)
		VALUES(?,?);
	`
	res, err := r.db.Exec(q, i.URL, i.Title)
	if err != nil {
		return nil, ErrInsertError
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrInsertError
	}
	i.ID = id
	return i, nil
}

func (r *ImageSqliteRepository) Get(id int64) (*model.Image, error) {
	i := model.Image{}
	err := r.db.Get(&i, "SELECT id, COALESCE(title, ''), COALESCE(url, '') FROM images WHERE id=?", id)
	if err != nil {
		return nil, ErrInsertError
	}
	return &i, nil
}

func (r *ImageSqliteRepository) Migrate() error {
	q := `
	CREATE TABLE IF NOT EXISTS images(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL DEFAULT '',
		url TEXT NOT NULL DEFAULT ''
	);
	`
	_, err := r.db.Exec(q)
	if err != nil {
		return ErrMigrationError
	}

	q1 := `
	CREATE TABLE IF NOT EXISTS article_images(
		article_id int64,
		image_id int64,
		PRIMARY KEY(article_id, image_id),
		CONSTRAINT fk_images
			FOREIGN KEY(image_id)
			REFERENCES images(id),
		CONSTRAINT fk_articles
			FOREIGN KEY(article_id)
			REFERENCES articles(id)
	);
	`
	_, err = r.db.Exec(q1)
	if err != nil {
		return ErrMigrationError
	}

	q2 := `
	CREATE TABLE IF NOT EXISTS newsfeed_images(
		newsfeed_id int64,
		image_id int64,
		PRIMARY KEY(newsfeed_id, image_id),
		CONSTRAINT fk_images
			FOREIGN KEY(image_id)
			REFERENCES images(id),
		CONSTRAINT fk_newsfeeds
			FOREIGN KEY(newsfeed_id)
			REFERENCES newsfeeds(id)
	);
	`
	_, err = r.db.Exec(q2)
	if err != nil {
		return ErrMigrationError
	}

	return nil
}
