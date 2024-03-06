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

func (s *SubscriptionService) Subscribe(f *gofeed.Feed, userId int64) (*model.Newsfeed, error) {
	// transform gofeed.Feed into model.Newsfeed
	nf, err := model.FeedFromRemote(f)
	if err != nil {
		fmt.Println("error converting feed: ", err)
		return nil, ErrInternalFailure
	}

	// transform Feed.Items into Newsfeed.Articles
	articles := []*model.Article{}
	for _, item := range f.Items {
		article, err := model.ArticleFromRemote(item)
		if err != nil {
			fmt.Println("error converting article: ", err)
			return nil, ErrInternalFailure
		}
		articles = append(articles, article)
	}
	// add articles to Newsfeed
	nf.Articles = articles
	feed, err := s.subRepo.AddAggregateSubscription(nf, userId)
	if err != nil {
		fmt.Println("error committing transaction: ", err)
		return nil, err
	}

	// return subscription model
	return feed, nil
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

func (s *SubscriptionService) GetArticles(userId int64) ([]*model.Article, error) {
	articles, err := s.subRepo.FindArticles(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *SubscriptionService) GetNewsfeeds(userId int64) ([]*model.NewsfeedExtended, error) {
	feeds, err := s.subRepo.FindNewsfeeds(userId)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
