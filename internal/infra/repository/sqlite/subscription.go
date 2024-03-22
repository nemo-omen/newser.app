package sqlite

import (
	"fmt"
	"sort"

	"github.com/jmoiron/sqlx"
	"newser.app/internal/domain/entity"
	"newser.app/internal/dto"
	"newser.app/shared"
)

type SubscriptionSqliteRepo struct {
	db *sqlx.DB
}

func NewSubscriptionSqliteRepo(db *sqlx.DB) *SubscriptionSqliteRepo {
	return &SubscriptionSqliteRepo{
		db: db,
	}
}

func (r *SubscriptionSqliteRepo) Create(subscription *dto.SubscriptionDTO) error {
	return nil
}

func (r *SubscriptionSqliteRepo) Delete(subscriptionID string) error {
	return nil
}

func (r *SubscriptionSqliteRepo) GetNewsfeed(userID, feedID string) (*dto.NewsfeedDTO, error) {
	feed := &dto.NewsfeedDTO{}
	err := r.db.Get(
		feed, `
			SELECT
				newsfeeds.*,
			COALESCE(images.title, '') as feed_image_title,
			COALESCE(images.url, '') as feed_image_url
			FROM
				subscriptions
			LEFT JOIN newsfeeds ON subscriptions.newsfeed_id = newsfeeds.id
			LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
			LEFT JOIN images ON newsfeed_images.image_id = images.id
			WHERE
				subscriptions.newsfeed_id = ?
			AND subscriptions.user_id = ?
			LIMIT 1;
		`,
		feedID,
		userID,
	)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to get subscribed newsfeed",
			"SubscriptionSqliteRepo.GetNewsfeed",
			"entity.Article",
		)
	}

	feedArticles := []*dto.ArticleDTO{}
	err = r.db.Select(
		&feedArticles, `
			SELECT * FROM articles WHERE newsfeed_id = ? ORDER BY published_parsed DESC LIMIT 10;`,
		feedID,
	)
	if err != nil || len(feedArticles) == 0 {
		return nil, shared.NewAppError(
			err,
			"Failed to get subscribed newsfeed articles",
			"SubscriptionSqliteRepo.GetNewsfeed",
			"entity.Article",
		)
	}

	for _, a := range feedArticles {
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
				"SubscriptionSqliteRepo.GetNewsfeed",
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
			a.ID,
		)
		if err != nil {
			a.Read = false
		} else {
			a.Read = true
		}

		savedCollection := dto.CollectionDTO{}
		err = r.db.Get(
			&savedCollection, `
			SELECT *
			FROM collections
			WHERE user_id = ? AND title = "saved";
		`,
			userID,
		)
		if err != nil {
			return nil, shared.NewAppError(
				err,
				"Failed to get saved collection",
				"SubscriptionSqliteRepo.GetNewsfeed",
				"entity.Collection",
			)
		}
		savedArticleId := ""
		err = r.db.Get(
			&savedArticleId, `
			SELECT article_id
			FROM collection_articles
			WHERE collection_id = ? AND article_id = ?;
		`,
			savedCollection.ID,
			a.ID,
		)
		if err != nil {
			a.Saved = false
		} else {
			a.Saved = true
		}
		a.FeedID = feedID
		a.FeedTitle = feed.Title
		a.FeedImageURL = feed.ImageURL
		a.FeedImageTitle = feed.ImageTitle
		a.FeedSlug = feed.Slug

	}
	feed.Articles = feedArticles
	return feed, nil
}

func (r *SubscriptionSqliteRepo) GetArticle(userID, articleID string) (*dto.ArticleDTO, error) {
	article := &dto.ArticleDTO{}
	articleQuery := `
	SELECT
		articles.*,
		newsfeeds.site_url as feed_site_url,
		newsfeeds.title as newsfeed_title,
		COALESCE(images.title, '') as feed_image_title,
		COALESCE(images.url, '') as feed_image_url
	FROM
		articles
		LEFT JOIN newsfeeds ON articles.newsfeed_id = newsfeeds.id
		LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
		LEFT JOIN images ON newsfeed_images.image_id = images.id
	WHERE articles.id = ?;
	`
	err := r.db.Get(article, articleQuery, articleID)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to get article",
			"SubscriptionSqliteRepo.GetArticle",
			"entity.Article",
		)
	}

	categories := []dto.CategoryDTO{}
	err = r.db.Select(
		&categories, `
		SELECT
			categories.*
		FROM 
			articles
			LEFT JOIN article_categories ON articles.id = article_categories.article_id
			LEFT JOIN categories ON article_categories.category_id = categories.id
		WHERE articles.id = ? AND categories.term != '';`,
		articleID)

	if err == nil {
		article.Categories = categories
	}

	readCollection := dto.CollectionDTO{}
	err = r.db.Get(
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
			"SubscriptionSqliteRepo.GetArticle",
			"entity.Collection",
		)
	}

	type CollectionArticle struct {
		ArticleID    string `db:"article_id"`
		CollectionId string `db:"collection_id"`
	}

	read := CollectionArticle{}

	err = r.db.Get(
		&read, `
		SELECT *
		FROM collection_articles
		WHERE collection_id = ? AND article_id = ?;`,
		readCollection.ID,
		articleID,
	)
	if err != nil {
		article.Read = false
	} else {
		article.Read = true
	}

	savedCollection := dto.CollectionDTO{}
	err = r.db.Get(
		&savedCollection, `
		SELECT *
		FROM collections
		WHERE user_id = ? AND title = "saved";
	`,
		userID,
	)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to get saved collection",
			"SubscriptionSqliteRepo.GetArticle",
			"entity.Collection",
		)
	}
	saved := CollectionArticle{}
	err = r.db.Get(
		&saved, `
		SELECT *
		FROM collection_articles
		WHERE collection_id = ? AND article_id = ?;`,
		savedCollection.ID,
		articleID,
	)
	if err != nil {
		article.Saved = false
	} else {
		article.Saved = true
	}

	return article, nil
}

func (r *SubscriptionSqliteRepo) Subscribe(userID string, feed dto.NewsfeedDTO) (*dto.SubscriptionDTO, error) {
	// check if feed exists
	storedFeed := dto.NewsfeedDTO{}
	err := r.db.Get(
		&storedFeed, `
		SELECT * FROM newsfeeds WHERE feed_url = ?;`,
		feed.FeedURL,
	)

	if err == nil {
		// feed exists, set ID & image ID
		// to the stored feed's values
		// so we can use it to insert the subscription
		// and avoid duplicate entries
		feed.ID = storedFeed.ID
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to start transaction",
			"SubscriptionSqliteRepo.Subscribe",
			"entity.Subscription",
		)
	}
	defer tx.Rollback()

	feedQuery := `
	INSERT INTO newsfeeds (id, title, site_url, feed_url, description, copyright, language, feed_type, slug)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err = tx.Exec(
		feedQuery,
		feed.ID,
		feed.Title,
		feed.SiteURL,
		feed.FeedURL,
		feed.Description,
		feed.Copyright,
		feed.Language,
		feed.FeedType,
		feed.Slug,
	)
	if err != nil {
		fmt.Println("err line 91: ", err)
		return nil, shared.NewAppError(
			err,
			"Failed to insert newsfeed",
			"SubscriptionSqliteRepo.Subscribe",
			"entity.Newsfeed",
		)
	}

	imgQuery := `
	INSERT INTO images (id, title, url)
	VALUES (?, ?, ?)
	ON CONFLICT (id) DO NOTHING;
	`
	img := entity.NewImage(feed.ImageURL, feed.ImageTitle)
	_, err = tx.Exec(imgQuery, img.ID, img.Title, img.URL)
	if err != nil {
		fmt.Println("err line 107: ", err)
		return nil, shared.NewAppError(
			err,
			"Failed to insert image",
			"SubscriptionSqliteRepo.Subscribe",
			"entity.Image",
		)
	}

	newsfeedImageQuery := `
	INSERT INTO newsfeed_images (newsfeed_id, image_id)
	VALUES (?, ?)
	ON CONFLICT (newsfeed_id, image_id) DO NOTHING;
	`
	_, err = tx.Exec(newsfeedImageQuery, feed.ID, img.ID)
	if err != nil {
		fmt.Println("err line 123: ", err)
		return nil, shared.NewAppError(
			err,
			"Failed to insert newsfeed image",
			"SubscriptionSqliteRepo.Subscribe",
			"entity.NewsfeedImage",
		)
	}

	for _, a := range feed.Articles {
		// check if article exists
		storedArticle := dto.ArticleDTO{}
		err := r.db.Get(
			&storedArticle, `
			SELECT * FROM articles WHERE article_link = ?;`,
			a.Link,
		)

		if err == nil {
			// article exists, set ID
			// to the stored article's values
			// so we can use it to insert the article
			// and avoid duplicate entries
			a.ID = storedArticle.ID
		}

		articleQuery := `
		INSERT INTO articles (id, title, description, content, article_link, published, published_parsed, updated, updated_parsed, guid, slug, newsfeed_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
		`
		_, err = tx.Exec(
			articleQuery,
			a.ID,
			a.Title,
			a.Description,
			a.Content,
			a.Link,
			a.Published,
			a.PublishedParsed,
			a.Updated,
			a.UpdatedParsed,
			a.GUID,
			a.Slug,
			feed.ID,
		)
		if err != nil {
			fmt.Println("err line 170: ", err)
			return nil, shared.NewAppError(
				err,
				"Failed to insert article",
				"SubscriptionSqliteRepo.Subscribe",
				"entity.Article",
			)
		}

		collection := dto.CollectionDTO{}
		err = r.db.Get(
			&collection, `
			SELECT *
			FROM collections
			WHERE user_id = ? AND title = "unread";
		`,
			userID,
		)
		if err != nil {
			return nil, shared.NewAppError(
				err,
				"Failed to get unread collection",
				"SubscriptionSqliteRepo.Subscribe",
				"entity.Collection",
			)
		}

		collectionArticleQuery := `
		INSERT INTO collection_articles (collection_id, article_id)
		VALUES (?, ?);
		`
		_, err = tx.Exec(collectionArticleQuery, collection.ID, a.ID)
		if err != nil {
			fmt.Println("err line 312: ", err)
			return nil, shared.NewAppError(
				err,
				"Failed to insert collection article",
				"SubscriptionSqliteRepo.Subscribe",
				"entity.CollectionArticle",
			)
		}

		for _, c := range a.Categories {
			storedCategory := dto.CategoryDTO{}
			err := r.db.Get(
				&storedCategory, `
				SELECT * FROM categories WHERE term = ?;`,
				c.Term,
			)
			if err == nil {
				c.ID = storedCategory.ID
			}

			categoryQuery := `
			INSERT INTO categories (id, term)
			VALUES (?, ?) ON CONFLICT (term) DO NOTHING;`
			_, err = tx.Exec(categoryQuery, c.ID, c.Term)
			if err == nil {
				articleCategoryQuery := `
			INSERT INTO article_categories (article_id, category_id)
			VALUES (?, ?);
			`
				_, err = tx.Exec(articleCategoryQuery, a.ID, c.ID)
				if err != nil {
					fmt.Println("err line 211: ", err)
					return nil, shared.NewAppError(
						err,
						"Failed to insert article category",
						"SubscriptionSqliteRepo.Subscribe",
						"entity.ArticleCategory",
					)
				}
			}
		}
	}

	subscriptionQuery := `
	INSERT INTO subscriptions (user_id, newsfeed_id)
	VALUES (?, ?);
	`
	_, err = tx.Exec(subscriptionQuery, userID, feed.ID)
	if err != nil {
		fmt.Println("err line 223: ", err)
		return nil, shared.NewAppError(
			err,
			"Failed to prepare subscription statement",
			"SubscriptionSqliteRepo.Subscribe",
			"entity.Subscription",
		)
	}
	subscription := &dto.SubscriptionDTO{
		UserID: userID,
		FeedID: feed.ID,
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("err line 223: ", err)
		return nil, shared.NewAppError(
			err,
			"Failed to commit transaction",
			"SubscriptionSqliteRepo.Subscribe",
			"entity.Subscription",
		)
	}

	return subscription, nil
}

func (r *SubscriptionSqliteRepo) GetAllArticles(userID string) ([]*dto.ArticleDTO, error) {
	feeds, err := r.GetAllFeeds(userID)
	if err != nil {
		return nil, err
	}

	articles := []*dto.ArticleDTO{}
	for _, f := range feeds {
		feedArticles := []*dto.ArticleDTO{}
		err := r.db.Select(
			&feedArticles, `
			SELECT
				articles.*,
				newsfeeds.title as newsfeed_title,
				newsfeeds.site_url as feed_site_url,
				newsfeeds.slug as feed_slug,
				COALESCE(images.title, '') as feed_image_title,
				COALESCE(images.url, '') as feed_image_url
			FROM
				newsfeeds
				LEFT JOIN articles ON newsfeeds.id = articles.newsfeed_id
				LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
				LEFT JOIN images ON newsfeed_images.image_id = images.id
			WHERE
				newsfeeds.id = ?
			ORDER BY articles.published_parsed DESC
			LIMIT 10;
		`, f.ID)
		if err != nil {
			return nil, shared.NewAppError(
				err,
				"Failed to get subscribed articles",
				"SubscriptionSqliteRepo.GetAllArticles",
				"entity.Article",
			)
		}
		for _, article := range feedArticles {
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
		}
		articles = append(articles, feedArticles...)
	}
	sort.SliceStable(articles, func(i, j int) bool {
		return articles[i].PublishedParsed.After(articles[j].PublishedParsed)
	})

	return articles, nil
}

func (r *SubscriptionSqliteRepo) GetAllFeeds(userID string) ([]*dto.NewsfeedDTO, error) {
	feeds := []*dto.NewsfeedDTO{}
	err := r.db.Select(
		&feeds, `
		SELECT
		newsfeeds.*,
		COALESCE(images.title, '') as feed_image_title,
		COALESCE(images.url, '') as feed_image_url
		FROM
			subscriptions
		LEFT JOIN newsfeeds ON subscriptions.newsfeed_id = newsfeeds.id
		LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
		LEFT JOIN images ON newsfeed_images.image_id = images.id
		WHERE
			subscriptions.user_id = ?
	`,
		userID,
	)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to get subscribed newsfeeds",
			"SubscriptionSqliteRepo.GetAllFeeds",
			"entity.Newsfeed",
		)
	}
	return feeds, nil
}

func (r *SubscriptionSqliteRepo) GetFeedsInfo(userId string) ([]*dto.FeedInfoDTO, error) {
	feedInfos := []*dto.FeedInfoDTO{}
	err := r.db.Select(
		// TODO: Update this query so we
		// can limit the number of articles
		// used for our unread count.
		&feedInfos,
		`SELECT
			newsfeeds.id as feed_id,
			newsfeeds.title as feed_title,
			COALESCE(images.title, '') as image_title,
			COALESCE(images.url, '') as image_url,
			COUNT (articles.id) as unread_count
		FROM
			collections
			LEFT JOIN collection_articles ON collection_id = collection_articles.collection_id
			LEFT JOIN articles ON collection_articles.article_id = articles.id
			LEFT JOIN newsfeeds ON articles.newsfeed_id = newsfeeds.id
			LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
			LEFT JOIN images ON newsfeed_images.image_id = images.id
		WHERE collections.title = 'unread' AND user_id = ?
		GROUP BY articles.newsfeed_id;`,
		userId,
	)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to get subscribed newsfeeds info",
			"SubscriptionSqliteRepo.GetFeedsInfo",
			"entity.FeedInfo",
		)
	}

	return feedInfos, nil
}
