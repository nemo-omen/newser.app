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

func (s *SubscriptionService) Subscribe(f *gofeed.Feed, userId int64) (*model.Subscription, error) {
	// transform gofeed.Feed into model.Newsfeed
	nf := model.FeedFromRemote(f)
	persistednf, err := s.feedRepo.Create(nf)
	if err != nil {
		return nil, err
	}

	// transform Feed.Items into Newsfeed.Articles
	articles := []*model.Article{}
	for _, item := range f.Items {
		article := model.ArticleFromRemote(item, persistednf.ID, persistednf.Title, persistednf.FeedUrl)
		article.FeedId = nf.ID
		article.FeedTitle = nf.Title
		article.FeedUrl = nf.FeedUrl
		articles = append(articles, &article)
	}
	persistedArticles, err := s.articleRepo.CreateMany(articles)
	if err != nil {
		return nil, err
	}
	nf.Articles = persistedArticles
	unreadColl, err := s.collectionRepo.FindByTitle("unread")
	if err != nil {
		fmt.Println("No unread collection??: ", err)
		return nil, err
	}
	err = s.collectionRepo.InsertManyCollectionItems(persistedArticles, unreadColl.Id)
	if err != nil {
		fmt.Println("error adding items to collection: ", err)
		return nil, err
	}

	sub := &model.Subscription{
		NewsfeedId: nf.ID,
		UserId:     userId,
	}

	sub, err = s.subRepo.Create(sub)
	if err != nil {
		return nil, err
	}

	// return subscription model
	return sub, nil
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
