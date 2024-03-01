package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type ArticleSqliteRepo struct {
	DB *sqlx.DB
}

func NewArticleSqliteRepo(db *sqlx.DB) *ArticleSqliteRepo {
	return &ArticleSqliteRepo{
		DB: db,
	}
}

func (r *ArticleSqliteRepo) Get(id uint) (*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) Create(a *model.Article) (*model.Article, error) {
	q := `
	INSERT INTO articles(
		title,
		description,
		content,
		article_link,
		author,
		published,
		published_parsed,
		updated,
		updated_parsed,
		image,
		guid,
		slug,
		feed_id,
		feed_title,
		feed_url
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`

	// cats := ""
	// if len(a.Categories) > 0 {
	// 	marshaledCats, err := json.Marshal(a.Categories)
	// 	if err != nil {
	// 		cats = ""
	// 	} else {
	// 		cats = string(marshaledCats)
	// 	}
	// }

	res, err := r.DB.Exec(
		q,
		a.Title,
		a.Description,
		a.Content,
		a.ArticleLink,
		a.Author,
		a.Published,
		a.PublishedParsed,
		a.Updated,
		a.UpdatedParsed,
		a.Image,
		a.GUID,
		a.Slug,
		a.FeedId,
		a.FeedTitle,
		a.FeedUrl,
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
	return a, nil
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

func (r *ArticleSqliteRepo) Update(n *model.Article) (*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) Delete(id uint) error {
	return nil
}

func (r *ArticleSqliteRepo) FindBySlug(slug string) (*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) ArticlesByCollection(collectionId int64) ([]*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) ArticlesByNewsfeed(feedId int64) ([]*model.Article, error) {
	return nil, nil
}

func (r *ArticleSqliteRepo) Migrate() error {
	qb := `
	CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		content TEXT,
		article_link TEXT NOT NULL,
		author TEXT,
		published TEXT NOT NULL,
		published_parsed DATETIME NOT NULL,
		updated TEXT NOT NULL,
		updated_parsed DATETIME NOT NULL,
		image TEXT,
		guid TEXT,
		slug TEXT NOT NULL,
		feed_id int NOT NULL,
		feed_title TEXT NOT NULL,
		feed_url TEXT NOT NULL,
		CONSTRAINT fk_newsfeeds
			FOREIGN KEY (feed_id)
			REFERENCES newsfeeds(id)
	);
	`
	_, err := r.DB.Exec(qb)
	if err != nil {
		fmt.Println("error migrating articles: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating articles")
	}
	return nil
}
