package repository

import (
	"newser.app/infra/dto"
	"newser.app/model"
)

type UserRepository interface {
	Get(id int64) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	GetHashedPasswordByEmail(email string) (string, error)
	All() []*model.User
	Create(udto *dto.UserDTO) (*model.User, error)
	Update(udto *dto.UserDTO) (*model.User, error)
	Delete(id int64) error
	Migrate() error
}

type NewsfeedRepository interface {
	Get(id uint) (*model.Newsfeed, error)
	Create(n *model.Newsfeed) (*model.Newsfeed, error)
	Update(n *model.Newsfeed) (*model.Newsfeed, error)
	Delete(id uint) error
	FindBySlug(slug string) (*model.Newsfeed, error)
	Migrate() error
}

type ArticleRepository interface {
	Get(id uint) (*model.Article, error)
	Create(n *model.Article) (*model.Article, error)
	CreateMany(n []*model.Article) ([]*model.Article, error)
	Update(n *model.Article) (*model.Article, error)
	Delete(id uint) error
	FindBySlug(slug string) (*model.Article, error)
	ArticlesByCollection(collectionId int64) ([]*model.Article, error)
	ArticlesByNewsfeed(feedId int64) ([]*model.Article, error)
	Migrate() error
}

type SubscriptionRepository interface {
	Get(id int64) (*model.Subscription, error)
	Create(*model.Subscription) (*model.Subscription, error)
	All(userId int64) ([]model.Subscription, error)
	Update(*model.Subscription) (*model.Subscription, error)
	Delete(id int64) error
	FindBySlug(slug string) (*model.Subscription, error)
	Migrate() error
}

type CollectionRepository interface {
	Get(id int64) (*model.Collection, error)
	Create(*model.Collection) (*model.Collection, error)
	All(userId int64) ([]*model.Collection, error)
	Update(*model.Collection) (*model.Collection, error)
	Delete(id int64) error
	FindBySlug(slug string) (*model.Collection, error)
	FindByTitle(title string) (*model.Collection, error)
	InsertCollectionItem(itemId int64, collectionId int64) error
	InsertManyCollectionItems(aa []*model.Article, cId int64) error
	Migrate() error
}
