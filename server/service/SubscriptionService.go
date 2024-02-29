package service

import (
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

func (s SubscriptionService) Subscribe(f *gofeed.Feed, userId int64) (model.Subscription, error) {
	// transform gofeed.Feed into model.Newsfeed
	nf := model.FeedFromRemote(*f)
	nf, err := s.feedRepo.Create(nf)
	if err != nil {
		return model.Subscription{}, err
	}

	// transform Feed.Items into Newsfeed.Articles
	articles := []model.Article{}
	for _, item := range f.Items {
		article := model.ArticleFromRemote(item, nf.ID, nf.Title, nf.FeedUrl)
		article, err = s.articleRepo.Create(article)
		if err != nil {
			return model.Subscription{}, err
		}
		collection, err := s.collectionRepo.FindByTitle("unread")
		if err != nil {
			return model.Subscription{}, err
		}
		err = s.collectionRepo.InsertCollectionItem(article.ID, collection.Id)
		if err != nil {
			return model.Subscription{}, err
		}
		articles = append(articles, article)
	}
	nf.Articles = articles

	sub := model.Subscription{
		NewsfeedId: nf.ID,
		UserId:     userId,
	}

	sub, err = s.subRepo.Create(sub)
	if err != nil {
		return model.Subscription{}, err
	}

	// return subscription model
	return sub, nil
}

func (s SubscriptionService) Unsubscribe(feedId, userId uint) error {
	// return err if failure only
	return nil
}

func (s SubscriptionService) All(userId int64) ([]model.Subscription, error) {
	ss, err := s.subRepo.All(userId)
	if err != nil {
		return ss, err
	}
	return ss, nil
}

func (s SubscriptionService) Get(subscriptionId uint) (model.Subscription, error) {
	return model.Subscription{}, nil
}
