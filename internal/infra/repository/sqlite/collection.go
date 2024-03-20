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
	return nil, nil
}
func (r *CollectionSqliteRepo) GetBySlug(slug, userId string) (*dto.CollectionDTO, error) {
	return nil, nil
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
	return nil, nil
}
func (r *CollectionSqliteRepo) GetCollectionNewsfeeds(collectionID string) ([]*dto.NewsfeedDTO, error) {
	return nil, nil
}
