package sqlite

import (
	"github.com/jmoiron/sqlx"
	"newser.app/internal/dto"
)

type CollectionSqliteRepo struct {
	db *sqlx.DB
}

func NewCollectionSqliteRepo(db *sqlx.DB) *CollectionSqliteRepo {
	return &CollectionSqliteRepo{
		db: db,
	}
}

func (r *CollectionSqliteRepo) Create(collection *dto.CollectionDTO) error {
	return nil
}

func (r *CollectionSqliteRepo) Delete(collectionID string) error {
	return nil
}
func (r *CollectionSqliteRepo) Get(collectionID string) (*dto.CollectionDTO, error) {
	return nil, nil
}
func (r *CollectionSqliteRepo) GetByTitle(title, userId string) (*dto.CollectionDTO, error) {
	coll := &dto.CollectionDTO{}
	err := r.db.Get(coll, "SELECT * FROM collections WHERE title = ? AND user_id = ?", title, userId)
	if err != nil {
		return nil, err
	}
	return coll, nil
}
func (r *CollectionSqliteRepo) GetBySlug(slug, userId string) (*dto.CollectionDTO, error) {
	coll := &dto.CollectionDTO{}
	err := r.db.Get(coll, "SELECT * FROM collections WHERE slug = ? AND user_id = ?", slug, userId)
	if err != nil {
		return nil, err
	}
	return coll, nil
}
func (r *CollectionSqliteRepo) All(userID string) ([]*dto.CollectionDTO, error) {
	return nil, nil
}
func (r *CollectionSqliteRepo) AddArticle(collectionID, articleID string) error {
	return nil
}
func (r *CollectionSqliteRepo) RemoveArticle(collectionID, articleID string) error {
	return nil
}
func (r *CollectionSqliteRepo) AddNewsfeed(collectionID, newsfeedID string) error {
	return nil
}
func (r *CollectionSqliteRepo) RemoveNewsfeed(collectionID, newsfeedID string) error {
	return nil
}
func (r *CollectionSqliteRepo) GetCollectionArticles(collectionID string) ([]*dto.ArticleDTO, error) {
	articles := []*dto.ArticleDTO{}
	err := r.db.Select(
		&articles,
		`SELECT
			articles.*,
			newsfeeds.title as newsfeed_title,
			newsfeeds.site_url as feed_site_url,
			newsfeeds.slug as feed_slug,
			COALESCE(images.title, '') as feed_image_title,
			COALESCE(images.url, '') as feed_image_url
		FROM
			collections
			JOIN collection_articles ON collections.id = collection_articles.collection_id
			JOIN articles ON collection_articles.article_id = articles.id
			JOIN newsfeeds ON articles.newsfeed_id = newsfeeds.id
			JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
			JOIN images ON newsfeed_images.image_id = images.id
		WHERE collections.id = ? ORDER BY articles.published_parsed DESC;`,
		collectionID,
	)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *CollectionSqliteRepo) GetCollectionNewsfeeds(collectionID string) ([]*dto.NewsfeedDTO, error) {
	return nil, nil
}
