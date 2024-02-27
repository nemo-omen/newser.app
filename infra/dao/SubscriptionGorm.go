package dao

type SubscriptionGorm struct {
	UserId     uint
	User       UserGorm
	NewsfeedId uint
	Newsfeed   NewsfeedGorm
	Slug       string
}

func (SubscriptionGorm) TableName() string {
	return "subscriptions"
}
