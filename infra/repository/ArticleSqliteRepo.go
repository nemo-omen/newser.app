package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type ArticleSqliteRepo struct {
	db *sqlx.DB
}

func NewArticleSqliteRepo(db *sqlx.DB) *ArticleSqliteRepo {
	return &ArticleSqliteRepo{
		db: db,
	}
}

func (r *ArticleSqliteRepo) Get(id int64) (*model.Article, error) {
	a := &model.Article{}
	// TODO: Nice join query so we can get more info about article
	// like images, author, feed title, etc.
	q := `
	SELECT
		articles.*,
		newsfeeds.id as feed_id,
		newsfeeds.title as feed_title,
		newsfeeds.feed_url as feed_url,
		newsfeeds.site_url as feed_site_url,
		newsfeeds.slug as feed_slug,
		COALESCE(people.name, '') as name,
		COALESCE(people.email, '') as email,
		COALESCE(images.title, '') as feed_image_title,
    	COALESCE(images.url, '') as feed_image_url
	FROM
		articles
		LEFT JOIN newsfeeds ON articles.feed_id = newsfeeds.id 
		LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
		LEFT JOIN images ON newsfeed_images.image_id = images.id
		LEFT JOIN article_people ON articles.id = article_people.article_id
		LEFT JOIN people ON article_people.person_id = people.id 
	WHERE
		articles.id = ?
	`
	err := r.db.Get(a, q, id)
	if err != nil {
		fmt.Println("error getting article: ", err.Error())
		return nil, err
	}
	return a, nil
}

func (r *ArticleSqliteRepo) Create(a *model.Article) (*model.Article, error) {
	art, err := InsertArticle(r.db, a)
	if err != nil {
		return nil, err
	}
	return art, nil
}

func (r *ArticleSqliteRepo) CreateMany(aa []*model.Article) ([]*model.Article, error) {
	allPersisted := []*model.Article{}
	for _, a := range aa {
		persisted, err := r.Create(a)

		if err != nil {
			return nil, err
		}
		allPersisted = append(allPersisted, persisted)
	}
	return allPersisted, nil
}

func (r *ArticleSqliteRepo) Update(a *model.Article) (*model.Article, error) {
	stmt, err := r.db.PrepareNamed(`
	UPDATE articles
		SET id=:id,
			title=:title,
			description=:description,
			content=:content,
			article_link=:article_link,
			published=:published,
			published_parsed=:published_parsed,
			updated=:updated,
			updated_parsed=:updated_parsed,
			guid=:guid,
			slug=:slug,
			feed_id=:feed_id,
			read=:read
	`)

	if err != nil {
		fmt.Println("error preparing update stmt: ", err.Error())
		return nil, err
	}

	_, err = stmt.Exec(a)
	if err != nil {
		fmt.Println("error updating article: ", err.Error())
		return nil, err
	}
	return a, nil
}

func (r *ArticleSqliteRepo) Delete(id int64) error {
	return nil
}

func (r *ArticleSqliteRepo) FindBySlug(slug string) (*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) ArticlesByCollection(collectionId int64) ([]*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) ArticlesByNewsfeed(feedId int64) ([]*model.Article, error) {
	aa := []*model.Article{}
	q := `
	SELECT
		newsfeeds.id as feed_id,
		newsfeeds.title as feed_title,
		newsfeeds.feed_url as feed_url,
		newsfeeds.site_url as feed_site_url,
		newsfeeds.slug as feed_slug,
		articles.*,
		COALESCE(images.title, '') as feed_image_title,
    	COALESCE(images.url, '') as feed_image_url
	FROM
		newsfeeds
		LEFT JOIN articles ON newsfeeds.id = articles.feed_id
		LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
		LEFT JOIN images ON newsfeed_images.image_id = images.id
	WHERE
		newsfeeds.id = ?
	ORDER BY articles.published_parsed DESC
	LIMIT 10;
	`
	err := r.db.Select(&aa, q, feedId)
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}
	return aa, nil
}

func (r *ArticleSqliteRepo) Migrate() error {
	qb := `
	CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		content TEXT,
		article_link TEXT NOT NULL,
		published TEXT NOT NULL,
		published_parsed DATETIME NOT NULL,
		updated TEXT NOT NULL,
		updated_parsed DATETIME NOT NULL,
		guid TEXT,
		slug TEXT NOT NULL,
		feed_id int NOT NULL,
		read BOOLEAN NOT NULL,
		CONSTRAINT fk_newsfeeds
			FOREIGN KEY (feed_id)
			REFERENCES newsfeeds(id)
	);
	`
	_, err := r.db.Exec(qb)
	if err != nil {
		fmt.Println("error migrating articles: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating articles")
	}
	return nil
}
