package sqlite

import (
	"github.com/jmoiron/sqlx"
	"newser.app/internal/dto"
	"newser.app/shared"
)

type NewsfeedSqliteRepo struct {
	db *sqlx.DB
}

func NewNewsfeedSqliteRepo(db *sqlx.DB) *NewsfeedSqliteRepo {
	return &NewsfeedSqliteRepo{
		db: db,
	}
}

func (r *NewsfeedSqliteRepo) InsertArticle(article *dto.ArticleDTO) error {
	stmt := `
	INSERT INTO articles (
		id,
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
		newsfeed_id
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	) ON CONFLICT (article_link, guid) DO UPDATE SET
		title = excluded.title,
		description = excluded.description,
		content = excluded.content,
		published = excluded.published,
		published_parsed = excluded.published_parsed,
		updated = excluded.updated,
		updated_parsed = excluded.updated_parsed,
		slug = excluded.slug;
	`
	_, err := r.db.Exec(stmt,
		article.ID,
		article.Title,
		article.Description,
		article.Content,
		article.Link,
		article.Published,
		article.PublishedParsed,
		article.Updated,
		article.UpdatedParsed,
		article.GUID,
		article.Slug,
		article.FeedID,
	)
	if err != nil {
		return shared.NewAppError(
			err,
			"Failed to insert article",
			"NewsfeedSqliteRepo.InsertArticle",
			"entity.Article",
		)
	}
	return nil
}
