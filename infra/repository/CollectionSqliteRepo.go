package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type CollectionSqliteRepo struct {
	db *sqlx.DB
}

func (r *CollectionSqliteRepo) Migrate() error {
	q := `
	CREATE TABLE IF NOT EXISTS collections(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		slug TEXT NOT NULL,
		user_id INT NOT NULL,
		CONSTRAINT fk_users
		FOREIGN KEY (user_id)
		REFERENCES users (id)
		ON DELETE CASCADE
	)
	`
	fmt.Println("Migrating collections table...")
	_, err := r.db.Exec(q)
	if err != nil {
		fmt.Println("error migrating collections: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating collections")
	}

	qa := `
	CREATE TABLE IF NOT EXISTS collection_articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		article_id INTEGER NOT NULL,
		collection_id INTEGER NOT NULL,
		CONSTRAINT fk_articles
			FOREIGN KEY (article_id)
				REFERENCES articles(id)
				ON DELETE CASCADE,
		CONSTRAINT fk_collections
			FOREIGN KEY(collection_id)
				REFERENCES collections(id)
				ON DELETE CASCADE
	);
	`
	fmt.Println("Migrating collection_articles table...")
	_, err = r.db.Exec(qa)
	if err != nil {
		fmt.Println("error migrating collections_articles: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating collection_articles")
	}
	return err
}

func NewCollectionSqliteRepo(db *sqlx.DB) *CollectionSqliteRepo {
	return &CollectionSqliteRepo{db: db}
}

func (r *CollectionSqliteRepo) Get(id int64) (*model.Collection, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) Create(c *model.Collection) (*model.Collection, error) {
	q := `
	INSERT INTO collections(title, slug, user_id)
		VALUES(?, ?, ?);
	`
	res, err := r.db.Exec(q, c.Title, c.Slug, c.UserId)

	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	c.Id = id
	return c, nil
}

func (r *CollectionSqliteRepo) All(userId int64) ([]*model.Collection, error) {
	ss := []*model.Collection{}
	err := r.db.Select(&ss, "SELECT * FROM collections WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (r *CollectionSqliteRepo) Update(s *model.Collection) (*model.Collection, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) Delete(id int64) error {
	return nil
}

func (r *CollectionSqliteRepo) FindByTitle(title string) (*model.Collection, error) {
	coll := &model.Collection{}
	err := r.db.Get(coll, "SELECT * FROM collections WHERE title=?", title)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

func (r *CollectionSqliteRepo) FindBySlug(slug string) (*model.Collection, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) InsertCollectionItem(articleId int64, collectionId int64) error {
	q := `
	INSERT INTO collection_articles(article_id, collection_id)
		VALUES(?, ?)
	`
	_, err := r.db.Exec(q, articleId, collectionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CollectionSqliteRepo) DeleteCollectionItem(itemId, collectionId int64) error {
	q := `DELETE FROM collection_articles WHERE article_id=? AND collection_id=?`
	_, err := r.db.Exec(q, itemId, collectionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CollectionSqliteRepo) InsertManyCollectionItems(aa []*model.Article, cId int64) error {
	for _, a := range aa {
		err := r.InsertCollectionItem(a.ID, cId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CollectionSqliteRepo) GetArticles(collectionId, userId int64) ([]*model.Article, error) {
	collectionArticles := []*model.Article{}
	err := r.db.Select(&collectionArticles, `
	SELECT
		articles.*,
		newsfeeds.title as feed_title,
		newsfeeds.feed_url as feed_url,
		newsfeeds.site_url as feed_site_url,
		newsfeeds.slug as feed_slug,
		COALESCE(images.url, '') as feed_image_url,
		COALESCE(images.title, '') as feed_image_title
		FROM
			collections
		LEFT JOIN collection_articles ON collections.id = collection_articles.collection_id
		LEFT JOIN articles ON collection_articles.article_id = articles.id
		LEFT JOIN newsfeeds ON articles.feed_id = newsfeeds.id
		LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
		LEFT JOIN images ON newsfeed_images.image_id = images.id
		WHERE
			collections.id = ?
		AND
			collections.user_id = ?
		ORDER BY articles.published_parsed DESC;
	`, collectionId, userId)
	if err != nil {
		return nil, err
	}
	return collectionArticles, nil
}

func (r *CollectionSqliteRepo) GetArticlesByCollectionName(collectionName string, userId int64) ([]*model.Article, error) {
	collectionArticles := []*model.Article{}
	err := r.db.Select(&collectionArticles, `
	SELECT
		articles.*,
		newsfeeds.title as feed_title,
		newsfeeds.feed_url as feed_url,
		newsfeeds.site_url as feed_site_url,
		newsfeeds.slug as feed_slug,
		COALESCE(images.url, '') as feed_image_url,
		COALESCE(images.title, '') as feed_image_title
		FROM
			collections
		LEFT JOIN collection_articles ON collections.id = collection_articles.collection_id
		LEFT JOIN articles ON collection_articles.article_id = articles.id
		LEFT JOIN newsfeeds ON articles.feed_id = newsfeeds.id
		LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
		LEFT JOIN images ON newsfeed_images.image_id = images.id
		WHERE
			collections.title = ?
		AND
			collections.user_id = ?
		ORDER BY articles.published_parsed DESC;
	`, collectionName, userId)
	if err != nil {
		return nil, err
	}
	return collectionArticles, nil
}

func (r *CollectionSqliteRepo) GetFeeds(collectionId, userId int64) ([]*model.NewsfeedExtended, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) GetFeedsByCollectionName(collectionName string, userId int64) ([]*model.NewsfeedExtended, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) MarkArticleRead(articleId, userId int64) error {
	fmt.Println("STARTING TX")
	tx := r.db.MustBegin()
	defer tx.Rollback()

	read := model.Collection{}
	err := tx.Get(&read, "SELECT * FROM collections WHERE title='read' AND user_id=? LIMIT 1", userId)
	if err != nil {
		fmt.Println("error selecting read collection: ", err.Error())
		return ErrNotFound
	}

	unread := model.Collection{}
	err = tx.Get(&unread, "SELECT * FROM collections WHERE title='unread' AND user_id=? LIMIT 1", userId)
	if err != nil {
		fmt.Println("error selecting unread collection: ", err.Error())
		return ErrNotFound
	}

	insertQ := `
	INSERT INTO collection_articles(article_id, collection_id)
	VALUES(?, ?)
	`
	_, err = tx.Exec(insertQ, articleId, read.Id)
	if err != nil {
		fmt.Println("error inserting into 'read' collection: ", err.Error())
		return ErrInsertError
	}

	deleteQ := `DELETE FROM collection_articles WHERE article_id=? AND collection_id=?`
	_, err = tx.Exec(deleteQ, articleId, unread.Id)
	if err != nil {
		fmt.Println("error removing from 'unread' collection: ", err.Error())
		return ErrInsertError
	}

	var storedArticle model.Article
	err = tx.Get(&storedArticle, "SELECT * FROM articles WHERE id=?", articleId)
	if err != nil {
		fmt.Println("error selecting article: ", err.Error())
		return ErrNotFound
	}

	storedArticle.Read = true

	stmt, err := tx.PrepareNamed(`
	UPDATE articles
		SET title=:title,
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
		WHERE id=:id
	`)

	if err != nil {
		fmt.Println("error preparing article update stmt: ", err.Error())
		return ErrInsertError
	}

	_, err = stmt.Exec(&storedArticle)

	if err != nil {
		fmt.Println("error updating article: ", err.Error())
		return ErrInsertError
	}

	fmt.Println("COMMITTING TX")
	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error: ", err)
		return ErrTransactionError
	}
	return nil
}

func (r *CollectionSqliteRepo) MarkArticleUnread(articleId, userId int64) error {
	fmt.Println("STARTING TX")
	tx := r.db.MustBegin()
	defer tx.Rollback()

	read := model.Collection{}
	err := tx.Get(&read, "SELECT * FROM collections WHERE title='read' AND user_id=? LIMIT 1", userId)
	if err != nil {
		fmt.Println("error selecting read collection: ", err.Error())
		return ErrNotFound
	}

	unread := model.Collection{}
	err = tx.Get(&unread, "SELECT * FROM collections WHERE title='unread' AND user_id=? LIMIT 1", userId)
	if err != nil {
		fmt.Println("error selecting unread collection: ", err.Error())
		return ErrNotFound
	}

	insertQ := `
	INSERT INTO collection_articles(article_id, collection_id)
	VALUES(?, ?)
	`
	_, err = tx.Exec(insertQ, articleId, unread.Id)
	if err != nil {
		fmt.Println("error inserting into 'unread' collection: ", err.Error())
		return ErrInsertError
	}

	deleteQ := `DELETE FROM collection_articles WHERE article_id=? AND collection_id=?`
	_, err = tx.Exec(deleteQ, articleId, read.Id)
	if err != nil {
		fmt.Println("error removing from 'read' collection: ", err.Error())
		return ErrInsertError
	}

	var storedArticle model.Article
	err = tx.Get(&storedArticle, "SELECT * FROM articles WHERE id=?", articleId)
	if err != nil {
		fmt.Println("error selecting article: ", err.Error())
		return ErrNotFound
	}

	storedArticle.Read = false

	stmt, err := tx.PrepareNamed(`
	UPDATE articles
		SET title=:title,
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
		WHERE id=:id
	`)

	if err != nil {
		fmt.Println("error preparing article update stmt: ", err.Error())
		return ErrInsertError
	}

	_, err = stmt.Exec(&storedArticle)

	if err != nil {
		fmt.Println("error updating article: ", err.Error())
		return ErrInsertError
	}

	fmt.Println("COMMITTING TX")
	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error: ", err)
		return ErrTransactionError
	}
	return nil
}
