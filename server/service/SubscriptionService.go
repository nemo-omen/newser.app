package service

import (
	"fmt"

	"github.com/mmcdole/gofeed"
	"newser.app/infra/repository"
	"newser.app/model"
)

type SubscriptionService struct {
	subRepo        repository.SubscriptionRepository
	feedRepo       repository.NewsfeedRepository
	articleRepo    repository.ArticleRepository
	collectionRepo repository.CollectionRepository
}

func NewSubscriptionService(
	sr repository.SubscriptionRepository,
	fr repository.NewsfeedRepository,
	ar repository.ArticleRepository,
	cr repository.CollectionRepository,
) SubscriptionService {
	return SubscriptionService{subRepo: sr, feedRepo: fr, articleRepo: ar, collectionRepo: cr}
}

func (s *SubscriptionService) Subscribe(f *gofeed.Feed, userId int64) error {
	// transform gofeed.Feed into model.Newsfeed
	nf, err := model.FeedFromRemote(f)
	if err != nil {
		fmt.Println("error converting feed: ", err)
		return ErrInternalFailure
	}

	// transform Feed.Items into Newsfeed.Articles
	articles := []*model.Article{}
	for _, item := range f.Items {
		article, err := model.ArticleFromRemote(item)
		if err != nil {
			fmt.Println("error converting article: ", err)
			return ErrInternalFailure
		}
		articles = append(articles, article)
	}
	// add articles to Newsfeed
	nf.Articles = articles
	err = s.subRepo.AddAggregateSubscription(nf, userId)
	if err != nil {
		fmt.Println("error committing transaction: ", err)
	}

	// return subscription model
	return nil
}

func (s *SubscriptionService) Unsubscribe(feedId, userId uint) error {
	// return err if failure only
	return nil
}

func (s *SubscriptionService) All(userId int64) ([]model.Subscription, error) {
	ss, err := s.subRepo.All(userId)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (s *SubscriptionService) Get(subscriptionId uint) (*model.Subscription, error) {
	return nil, nil
}
