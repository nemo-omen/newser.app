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

func (r *ArticleSqliteRepo) Get(id uint) (*model.Article, error) {
	return nil, nil
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
	aa := []*model.Article{}
	err := r.db.Select(&aa, "SELECT * FROM articles WHERE feed_id=?", feedId)
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
