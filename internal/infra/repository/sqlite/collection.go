package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/internal/dto"
	"newser.app/shared"
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
	_, err := r.db.Exec(
		`INSERT INTO collection_articles (
			collection_id, article_id
			) VALUES (?, ?) ON CONFLICT (collection_id, article_id) DO NOTHING`,
		collectionID,
		articleID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *CollectionSqliteRepo) RemoveArticle(collectionID, articleID string) error {
	_, err := r.db.Exec("DELETE FROM collection_articles WHERE collection_id = ? AND article_id = ?", collectionID, articleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CollectionSqliteRepo) AddAndRemoveArticle(addCollectionID, removeCollectionID, articleID string) error {
	tx, err := r.db.Beginx()
	defer tx.Rollback()
	if err != nil {
		fmt.Println("Error starting transaction in AddAndRemoveArticle: ", err.Error())
		return shared.NewAppError(
			err,
			"Error starting transaction",
			"AddAndRemoveArticle",
			"CollectionSqliteRepo",
		)
	}
	_, err = tx.Exec(
		`INSERT INTO collection_articles (
			collection_id, article_id
		) VALUES (?, ?) 
			ON CONFLICT (
				collection_id, article_id
			) DO NOTHING`,
		addCollectionID,
		articleID,
	)
	if err != nil {
		fmt.Println("Error adding article to collection in AddAndRemoveArticle: ", err.Error())
		return shared.NewAppError(
			err,
			"Error adding article to collection",
			"AddAndRemoveArticle",
			"CollectionSqliteRepo",
		)
	}
	_, err = tx.Exec(
		`DELETE FROM collection_articles
		WHERE collection_id = ? AND article_id = ?`,
		removeCollectionID,
		articleID,
	)
	if err != nil {
		fmt.Println("Error removing article from collection in AddAndRemoveArticle: ", err.Error())
		return shared.NewAppError(
			err,
			"Error removing article from collection",
			"AddAndRemoveArticle",
			"CollectionSqliteRepo",
		)
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction in AddAndRemoveArticle: ", err.Error())
		return shared.NewAppError(
			err,
			"Error committing transaction",
			"AddAndRemoveArticle",
			"CollectionSqliteRepo",
		)
	}
	return nil
}

func (r *CollectionSqliteRepo) AddNewsfeed(collectionID, newsfeedID string) error {
	return nil
}
func (r *CollectionSqliteRepo) RemoveNewsfeed(collectionID, newsfeedID string) error {
	return nil
}
func (r *CollectionSqliteRepo) GetCollectionArticles(collectionID, userID string) ([]*dto.ArticleDTO, error) {
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
	for _, article := range articles {
		readCollection := dto.CollectionDTO{}
		err := r.db.Get(
			&readCollection, `
				SELECT *
				FROM collections
				WHERE user_id = ? AND title = "read";
			`,
			userID,
		)
		if err != nil {
			return nil, shared.NewAppError(
				err,
				"Failed to get read collection",
				"SubscriptionSqliteRepo.GetAllArticles",
				"entity.Collection",
			)
		}
		readArticleId := ""
		err = r.db.Get(
			&readArticleId, `
				SELECT article_id
				FROM collection_articles
				WHERE collection_id = ? AND article_id = ?;
			`,
			readCollection.ID,
			article.ID,
		)
		if err != nil {
			article.Read = false
		} else {
			article.Read = true
		}
		fmt.Println("read?: ", article.Read)
	}
	return articles, nil
}

func (r *CollectionSqliteRepo) GetCollectionNewsfeeds(collectionID string) ([]*dto.NewsfeedDTO, error) {
	return nil, nil
}
