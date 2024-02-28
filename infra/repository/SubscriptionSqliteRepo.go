package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type SubscriptionSqliteRepo struct {
	DB *sqlx.DB
}

func (r SubscriptionSqliteRepo) Migrate() error {
	fmt.Println("Migrating subscriptions table...")
	q := `
	CREATE TABLE IF NOT EXISTS subscriptions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		slug TEXT NOT NULL,
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
	_, err := r.DB.Exec(q)
	if err != nil {
		fmt.Println("error migrating subscriptions: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating subscriptions")
	}
	return err
}

func NewSubscriptionSqliteRepo(db *sqlx.DB) SubscriptionSqliteRepo {
	return SubscriptionSqliteRepo{DB: db}
}

func (r SubscriptionSqliteRepo) Get(id int64) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionSqliteRepo) Create(s model.Subscription) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionSqliteRepo) All(userId int64) ([]model.Subscription, error) {
	ss := []model.Subscription{}
	err := r.DB.Select(&ss, "SELECT * FROM subscriptions WHERE user_id=?", userId)
	if err != nil {
		return ss, err
	}
	return ss, nil
}

func (r SubscriptionSqliteRepo) Update(s model.Subscription) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionSqliteRepo) Delete(id int64) error {
	return nil
}

func (r SubscriptionSqliteRepo) FindBySlug(slug string) (model.Subscription, error) {
	return model.Subscription{}, nil
}
