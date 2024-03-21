package sqlite

import (
	"github.com/jmoiron/sqlx"
	"newser.app/internal/dto"
)

type ArticleSqliteRepo struct {
	db *sqlx.DB
}

func NewArticleSqliteRepo(db *sqlx.DB) *ArticleSqliteRepo {
	return &ArticleSqliteRepo{
		db: db,
	}
}

func (r *ArticleSqliteRepo) Get(articleID string) (*dto.ArticleDTO, error) {
	return nil, nil
}
func (r *ArticleSqliteRepo) GetBySlug(slug string) (*dto.ArticleDTO, error) {
	article := &dto.ArticleDTO{}
	err := r.db.Get(article, "SELECT * FROM articles WHERE slug = ?", slug)
	if err != nil {
		return nil, err
	}
	return article, nil
}
func (r *ArticleSqliteRepo) GetByURL(url string) (*dto.ArticleDTO, error) {
	return nil, nil
}
