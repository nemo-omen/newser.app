package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
	"newser.app/shared/util"
)

type SqliteArticleRepo struct {
	db *sqlx.DB
}

func NewSqliteArticleRepo(db *sqlx.DB) *SqliteArticleRepo {
	return &SqliteArticleRepo{
		db: db,
	}
}

func (r *SqliteArticleRepo) Get(ctx context.Context, id int64) (*model.Article, error) {
	const query = `
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
	a := &model.Article{}
	err := r.db.GetContext(ctx, a, query, id)
	if err != nil {
		return nil, fmt.Errorf("error getting article: %w", err)
	}
	return a, nil
}

func (r *SqliteArticleRepo) Create(ctx context.Context, a *model.Article) (*model.Article, error) {
	q := `
	INSERT INTO articles(
		title,
		description,
		content,
		article_link,
		published,
		published_parsed,
		updated,
		updated_parsed,
		guid,
		slug,
		feed_id,
		read
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);
	`
	res, err := db.Exec(
		q,
		a.Title,
		a.Description,
		a.Content,
		a.ArticleLink,
		a.Published,
		a.PublishedParsed,
		a.Updated,
		a.UpdatedParsed,
		a.GUID,
		util.Slugify(a.Title),
		a.FeedId,
		false,
	)
	if err != nil {
		fmt.Println("article insert err: ", err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	a.ID = id

	if a.Image != nil {
		_ = InsertArticleImageTx(db, a.Image, a.ID)
	}

	_ = InsertArticlePersonTx(db, &a.Person, a.ID)
	return a, nil
}
