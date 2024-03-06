package repository

import (
	"fmt"
	"sort"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type SubscriptionSqliteRepo struct {
	db *sqlx.DB
}

func (r *SubscriptionSqliteRepo) Migrate() error {
	fmt.Println("Migrating subscriptions table...")
	q := `
	CREATE TABLE IF NOT EXISTS subscriptions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INT NOT NULL,
		newsfeed_id INT NOT NULL,
		CONSTRAINT fk_users
			FOREIGN KEY (user_id)
				REFERENCES users (id),
		CONSTRAINT fk_newsfeeds
			FOREIGN KEY (newsfeed_id)
				REFERENCES newsfeeds (id)
	)
	`
	_, err := r.db.Exec(q)
	if err != nil {
		fmt.Println("error migrating subscriptions: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating subscriptions")
	}
	return nil
}

func NewSubscriptionSqliteRepo(db *sqlx.DB) *SubscriptionSqliteRepo {
	return &SubscriptionSqliteRepo{db: db}
}

func (r *SubscriptionSqliteRepo) Get(id int64) (*model.Subscription, error) {
	return nil, nil
}

func (r *SubscriptionSqliteRepo) Create(s *model.Subscription) (*model.Subscription, error) {
	q := `
	INSERT INTO subscriptions(user_id, newsfeed_id)
		VALUES(?, ?);
	`
	res, err := r.db.Exec(q, s.UserId, s.NewsfeedId)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	s.Id = id
	return s, nil
}

func (r *SubscriptionSqliteRepo) All(userId int64) ([]model.Subscription, error) {
	ss := []model.Subscription{}
	err := r.db.Select(&ss, "SELECT * FROM subscriptions WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (r *SubscriptionSqliteRepo) Update(s *model.Subscription) (*model.Subscription, error) {
	return nil, nil
}

func (r *SubscriptionSqliteRepo) Delete(id int64) error {
	return nil
}

func (r *SubscriptionSqliteRepo) FindBySlug(slug string) (*model.Subscription, error) {
	return nil, nil
}

func (r *SubscriptionSqliteRepo) AddAggregateSubscription(n *model.Newsfeed, userId int64) (*model.Newsfeed, error) {
	feed, err := InsertAggregateSubscriptionTx(r.db, n, userId)
	if err != nil {
		return nil, err
	}
	return feed, err
}

func (r *SubscriptionSqliteRepo) FindArticles(userId int64) ([]*model.Article, error) {
	feeds, err := r.FindNewsfeeds(userId)
	if err != nil {
		return nil, err
	}

	articles := []*model.Article{}
	for _, feed := range feeds {
		feedArticles := []*model.Article{}
		err := r.db.Select(&feedArticles, `
			SELECT
				newsfeeds.id as feed_id,
				newsfeeds.title as feed_title,
				newsfeeds.feed_url as feed_url,
				newsfeeds.site_url as feed_site_url,
				newsfeeds.slug as feed_slug,
				articles.*,
				COALESCE(images.title, '') as feed_image_title,
				COALESCE(images.url, '') as feed_image_url
			FROM
				newsfeeds
				LEFT JOIN articles ON newsfeeds.id = articles.feed_id
				LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
				LEFT JOIN images ON newsfeed_images.image_id = images.id
			WHERE
				newsfeeds.id = ?
			ORDER BY articles.published_parsed DESC
			LIMIT 10;
		`, feed.ID)

		if err != nil {
			return nil, err
		}
		articles = append(articles, feedArticles...)
	}
	sort.SliceStable(articles, func(i, j int) bool {
		return articles[i].PublishedParsed.After(articles[j].PublishedParsed)
	})
	return articles, nil
}

func (r *SubscriptionSqliteRepo) FindNewsfeeds(userId int64) ([]*model.NewsfeedExtended, error) {
	feedLinks := []*model.NewsfeedExtended{}
	err := r.db.Select(&feedLinks, `
	SELECT
		subscriptions.id as subscription_id,
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
	`, userId)

	if err != nil {
		fmt.Println("err getting feedstuff: ", err)
		return nil, err
	}

	return feedLinks, nil
}
