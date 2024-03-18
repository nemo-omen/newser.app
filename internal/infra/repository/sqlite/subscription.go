package sqlite

import (
	"github.com/jmoiron/sqlx"
	"newser.app/internal/dto"
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

func (r *SubscriptionSqliteRepo) GetAllArticles(userID string) ([]*dto.ArticleDTO, error) {
	feedArticles := []*dto.ArticleDTO{}
	// err := r.db.Select(
	// 	&feedArticles, `
	// 	-- awesome join here
	//  -- see SubscribedArticlesWithNewsfeeds query
	// 	`,
	// 	userID,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	return feedArticles, nil
}

func (r *SubscriptionSqliteRepo) GetAllFeeds(userID string) ([]*dto.NewsfeedDTO, error) {
	return nil, nil
}

func (r *SubscriptionSqliteRepo) GetFeedsInfo(feedID string) (*dto.FeedInfoDTO, error) {
	return nil, nil
}
